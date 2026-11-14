// Run clang only when producing the checked-in .syso, never when building Go.
_Thread_local int value = 7;

int tls_next(void) {
    return value++;
}
