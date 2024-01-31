use std::env;
use std::fs::{File, OpenOptions};
use std::io::{Read, Write};
use std::net::{TcpListener, UdpSocket};

use clap::Parser;
use log::{error, info};
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

fn main() {
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
        match File::create(TMPFILE) {
            Ok(_) => info!("✅ File write successful: {}", TMPFILE),
            Err(err) => error!("❌ Error creating file: {}", err),
        }
    }

    if args.file_read {
        match OpenOptions::new().read(true).open(TMPFILE) {
            Ok(mut file) => {
                let mut contents = String::new();
                if let Err(err) = file.read_to_string(&mut contents) {
                    error!("❌ Error reading file: {}", err);
                } else {
                    info!("✅ File read successful: {}", TMPFILE);
                }
            }
            Err(err) => error!("❌ Error reading file: {}", err),
        }
    }

    // Network operations
    if args.net_tcp {
        match TcpListener::bind("0.0.0.0:0") {
            // Dynamic port assignment
            Ok(listener) => {
                let addr = listener.local_addr().unwrap();
                info!("✅ TCP server spawned: {}", addr);

                // Listener would need to be handled, likely in a separate thread
            }
            Err(err) => error!("❌ Error starting TCP server: {}", err),
        }
    }

    if args.net_udp {
        match UdpSocket::bind("0.0.0.0:0") {
            Ok(socket) => {
                let addr = socket.local_addr().unwrap();
                info!("✅ UDP server spawned: {}", addr);

                // Socket would need to be handled for receiving/sending
            }
            Err(err) => error!("❌ Error starting UDP server: {}", err),
        }
    }

    if args.net_icmp {
        unimplemented!();
    }

    // Subprocess execution
    if let [program, program_args @ ..] = &args.subprocess[..] {
        info!("✴️ Spawning subproces...");
        match Command::new(program)
            .args(program_args)
            .env(
                "LOGPREFIX",
                format!("  {}", env::var(LOGPREFIX_ENV_VAR).unwrap_or_default()),
            )
            .spawn()
            .and_then(|mut child| child.wait())
        {
            Ok(code) => {
                info!(
                    "✅ {:?} terminated with exit code {}.",
                    args.subprocess, code
                );
            }
            Err(e) => {
                error!("❌ Subprocess failed: {}", e)
            }
        };
    }

    // Simulating crash with SIGKILL (Caution!)
    if args.crash {
        info!("🫡  Terminating with SIGKILL...");
        // This is potentially dangerous in a real-world Rust program.
        // Consider safer alternatives for controlled termination.
        unsafe { libc::kill(libc::getpid(), libc::SIGKILL) };
    }

    info!("⭐️ Success.");
}
