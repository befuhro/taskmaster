package task_master

//var signals = map[os.Signal]string {
//	syscall.SIGHUP: "", // SIGHUP
//	syscall.SIGHUP: "", // SIGINT
//	syscall.SIGHUP: "", // SIGQUIT
//	syscall.SIGHUP: "", // SIGILL
//	syscall.SIGHUP: "", // SIGTRAP
//	syscall.SIGHUP: "", // SIGABRT
//	syscall.SIGHUP: "", // SIGBUS
//	syscall.SIGHUP: "", // SIGFPE
//	syscall.SIGHUP: "", // SIGKILL
//	syscall.SIGHUP: "", // SIGUSR1
//	syscall.SIGHUP: "", // SIGSEGV
//	syscall.SIGHUP: "", // SIGUSR2
//	syscall.SIGHUP: "", // SIGPIPE
//	syscall.SIGHUP: "", // SIGALRM
//	syscall.SIGHUP: "", // SIGTERM
//	syscall.SIGHUP: "", // SIGSTKFLT
//	syscall.SIGHUP: "", // SIGCHLD
//	syscall.SIGHUP: "", // SIGCONT
//	syscall.SIGHUP: "", // SIGSTOP
//	syscall.SIGHUP: "", // SIGTSTP
//	syscall.SIGHUP: "", // SIGTTIN
//	syscall.SIGHUP: "", // SIGTTOU
//	syscall.SIGHUP: "", // SIGURG
//	syscall.SIGHUP: "", // SIGXCPU
//	syscall.SIGHUP: "", // SIGXFSZ
//	syscall.SIGHUP: "", // SIGVTALRM
//	syscall.SIGHUP: "", // SIGPROF
//	syscall.SIGHUP: "", // SIGWINCH
//	syscall.SIGHUP: "", // SIGIO
//	syscall.SIGHUP: "", // SIGPWR
//	syscall.SIGHUP: "", // SIGSYS
//}
//
//34) SIGRTMIN
//35) SIGRTMIN+1
//36) SIGRTMIN+2
//37) SIGRTMIN+3
//38) SIGRTMIN+4
//39) SIGRTMIN+5
//40) SIGRTMIN+6
//41) SIGRTMIN+7
//42) SIGRTMIN+8
//43) SIGRTMIN+9
//44) SIGRTMIN+10
//45) SIGRTMIN+11
//46) SIGRTMIN+12
//47) SIGRTMIN+13
//48) SIGRTMIN+14
//49) SIGRTMIN+15
//50) SIGRTMAX-14
//51) SIGRTMAX-13
//52) SIGRTMAX-12
//53) SIGRTMAX-11
//54) SIGRTMAX-10
//55) SIGRTMAX-9
//56) SIGRTMAX-8
//57) SIGRTMAX-7
//58) SIGRTMAX-6
//59) SIGRTMAX-5
//60) SIGRTMAX-4
//61) SIGRTMAX-3
//62) SIGRTMAX-2
//63) SIGRTMAX-1
//64) SIGRTMAX