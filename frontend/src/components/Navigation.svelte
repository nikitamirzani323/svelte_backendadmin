<script>
    async function handleLogout() {
        localStorage.clear();
        window.location.href = "/";
    }
</script>
<div class="sticky top-0 z-30 flex h-16 w-full justify-center bg-opacity-90 backdrop-blur transition-all duration-100 
    bg-[#1a73e8] text-base-content shadow-sm">
    <div class="navbar ">
        <div class="flex-none">
        <label for="my-drawer-2" class="btn btn-square btn-ghost drawer-button ">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="inline-block w-5 h-5 stroke-current"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"></path></svg>
        </label>
        </div>
        <div class="flex-1">
        <a class="btn btn-ghost normal-case text-xl text-white">SDSB</a>
        </div>
        <div class="flex-none">
        <button
            on:click={() => {
                handleLogout();
            }}
            class="btn btn-ghost btn-circle">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white fill-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
            </svg>
        </button>
        </div>
    </div>
</div>