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

    u32 mntns = get_mntns();
    u32 pid = bpf_get_current_pid_tgid() >> 32;
    u32 request = BPF_CORE_READ(ad, request);
    const char * name_ptr = BPF_CORE_READ(ad, name);

    bpf_printk("aa_audit! %s", name_ptr);

    char name[256];
    long name_len = 0;
    if(name_ptr) {
        name_len = bpf_probe_read_kernel_str(name, sizeof(name), name_ptr);
        if(name_len < 0) {
            return 0;
        }
    }

    struct bpf_dynptr event;
    u32 size = 4 + 4 + 4 + name_len;
    bpf_ringbuf_reserve_dynptr(&audit_log, size, 0, &event);
    bpf_dynptr_write(&event, 0, &mntns, 4, 0);
    bpf_dynptr_write(&event, 4, &pid, 4, 0);
    bpf_dynptr_write(&event, 8, &request, 4, 0);
    bpf_dynptr_write(&event, 12, &name, name_len, 0);
    bpf_ringbuf_submit_dynptr(&event, 0);

    return 0;
}
