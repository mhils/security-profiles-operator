use std::env;
use std::fs::OpenOptions;
use std::io::Write;
use std::net::{TcpListener, UdpSocket};

use anyhow::{bail, Context, Result};
use clap::Parser;
use log::info;
use std::process::Command;
use std::str;

use std::ffi::OsString;

const TMPFILE: &str = "/dev/null";
const LOGPREFIX_ENV_VAR: &str = "LOGPREFIX";

/// Demo program to exercise various syscalls.
#[derive(Parser, Debug)]
#[command()]
struct Args {
    /// Write a file.
    #[arg(long)]
    file_write: bool,

    /// Read a file.
    #[arg(long)]
    file_read: bool,

    /// Spawn a TCP socket.
    #[arg(long)]
    net_tcp: bool,

    /// Spawn a UDP socket.
    #[arg(long)]
    net_udp: bool,

    /// Spawn an ICMP socket.
    #[arg(long)]
    net_icmp: bool,

    /// Crash with SIGKILL.
    #[arg(long)]
    crash: bool,

    /// Execute a subprocess
    subprocess: Vec<OsString>,
}

fn main() -> Result<()> {
    env_logger::Builder::new()
        .format(|buf, record| {
            writeln!(
                buf,
                "{}[pid:{}] {}",
                env::var(LOGPREFIX_ENV_VAR).unwrap_or_default(),
                std::process::id(),
                record.args()
            )
        })
        .format_timestamp(None)
        .filter_level(log::LevelFilter::Info)
        .init();

    info!("⏩ {:?}", env::args().collect::<Vec<_>>()); // Logs the arguments

    let args: Args = Args::parse();

    // File operations
    if args.file_write {
        OpenOptions::new()
            .write(true)
            .open(TMPFILE)
            .context("❌ Failed to write file")?;
        info!("✅ File write successful: {}", TMPFILE);
    }

    if args.file_read {
        OpenOptions::new()
            .read(true)
            .open(TMPFILE)
            .context("❌ Failed to read file")?;
        info!("✅ File read successful: {}", TMPFILE);
    }

    // Network operations
    if args.net_tcp {
        let listener = TcpListener::bind("0.0.0.0:0").context("❌ Failed to start TCP server")?;
        info!("✅ TCP server spawned: {}", listener.local_addr()?);
    }

    if args.net_udp {
        let socket = UdpSocket::bind("0.0.0.0:0").context("❌ Failed to start UDP server")?;
        info!("✅ UDP server spawned: {}", socket.local_addr()?);
    }

    if args.net_icmp {
        unimplemented!();
    }

    // Subprocess execution
    if let [program, program_args @ ..] = &args.subprocess[..] {
        info!("✴️  Spawning {:?}...", args.subprocess);
        let exit_code = Command::new(program)
            .args(program_args)
            .env(
                "LOGPREFIX",
                format!("  {}", env::var(LOGPREFIX_ENV_VAR).unwrap_or_default()),
            )
            .spawn()
            .context("❌ Failed to spawn subprocess")?
            .wait()
            .context("❌ Subprocess failed")?;
        if exit_code.success() {
            info!(
                "✅ {:?} terminated with exit code {}.",
                args.subprocess, exit_code
            );
        } else {
            bail!("❌ {:?} terminated with exit code {}.",args.subprocess, exit_code);
        }
        
    }

    // Simulating crash with SIGKILL (Caution!)
    if args.crash {
        info!("🫡  Terminating with SIGKILL...");
        // This is potentially dangerous in a real-world Rust program.
        // Consider safer alternatives for controlled termination.
        unsafe { libc::kill(libc::getpid(), libc::SIGKILL) };
    }

    info!("⭐️ Success.");
    Ok(())
}
