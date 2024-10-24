//go:build none

// This file is manually included, to avoid CGo which would cause a circular
// import.

#include <stdint.h>
#include <sys/syscall.h>
#include <time.h>
#include <unistd.h>

#define FUTEX_WAIT    0
#define FUTEX_WAKE    1
#define FUTEX_PRIVATE 128

void tinygo_futex_wait(uint32_t *addr, uint32_t cmp) {
    syscall(SYS_futex, addr, FUTEX_WAIT|FUTEX_PRIVATE, cmp, NULL, NULL, 0);
}

void tinygo_futex_wait_timeout(uint32_t *addr, uint32_t cmp, uint64_t timeout) {
    struct timespec ts = {0};
    ts.tv_sec = timeout / 1000000000;
    ts.tv_nsec = timeout % 1000000000;
    syscall(SYS_futex, addr, FUTEX_WAIT|FUTEX_PRIVATE, cmp, &ts, NULL, 0);
}

void tinygo_futex_wake(uint32_t *addr, uint32_t num) {
    syscall(SYS_futex, addr, FUTEX_WAKE|FUTEX_PRIVATE, num, NULL, NULL, 0);
}
