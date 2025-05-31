#include <vmlinux.h>
#include <linux/limits.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>

char LICENSE[] SEC("license") = "Dual BSD/GPL";

#ifndef likely
#define likely(x) __builtin_expect((x), 1)
#endif

#define MAX_NAMESPACES 8096

struct {
	__uint(type, BPF_MAP_TYPE_RINGBUF);
	__uint(max_entries, 256 * 1024 /* 256 KB */);
} audit_log SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_LRU_HASH);
    __uint(max_entries, MAX_NAMESPACES);
    __type(key, u32);
    __type(value, u32);
} apparmor_violations SEC(".maps");

static __always_inline u32 get_mntns()
{
    struct task_struct * task = (struct task_struct *)bpf_get_current_task();
    return BPF_CORE_READ(task, nsproxy, mnt_ns, ns.inum);
}

SEC("kprobe/aa_audit")
int BPF_KPROBE(kprobe__aa_audit, int type, struct aa_profile * profile, struct apparmor_audit_data *ad)
{
    const int error = BPF_CORE_READ(ad, error);
    if (likely(!error)) {
        return 0;
    }
    if (type == AUDIT_APPARMOR_HINT || type == AUDIT_APPARMOR_STATUS) {
        return 0;
    }

    const char * name_ptr = BPF_CORE_READ(ad, name);

    bpf_printk("aa_audit! %s", name_ptr);

    // Update counter
    // XXX: We currently do not distinguish between complain and enforce mode.
    u64 cgroup_id = bpf_get_current_cgroup_id();
    u32 mntns = get_mntns();
    u32 count = 1;
    u32 *existing_count = bpf_map_lookup_elem(&apparmor_violations, &mntns);
    if (existing_count) {
        count += *existing_count;
    }
    bpf_map_update_elem(&apparmor_violations, &mntns, &count, BPF_ANY);

    if(!name_ptr) {
        return 0;
    }

    u32 request = BPF_CORE_READ(ad, request);
    char name[256];
    long name_len = bpf_probe_read_kernel_str(name, sizeof(name), name_ptr);
    if (name_len > 0) {
        struct bpf_dynptr event;
        u32 size = 4 + 4 + name_len;


        bpf_ringbuf_reserve_dynptr(&audit_log, size, 0, &event);
        bpf_dynptr_write(&event, 0, &mntns, 4, 0);
        bpf_dynptr_write(&event, 4, &request, 4, 0);
        bpf_dynptr_write(&event, 8, &name, name_len, 0);
        bpf_ringbuf_submit_dynptr(&event, 0);
    }

    return 0;
}
