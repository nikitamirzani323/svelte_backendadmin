<script>
  import { link } from "svelte-spa-router";
  let menu_master_child_flag = false
  let menu_togel_child_flag = false
  const togglemaster = () => {
		menu_master_child_flag = !menu_master_child_flag
	}
  const toggletogel = () => {
		menu_togel_child_flag = !menu_togel_child_flag
	}
</script>
<div class="drawer-side " style="scroll-behavior: smooth; scroll-padding-top: 5rem;">
    <label for="my-drawer-2" class="drawer-overlay"></label> 
    <ul class="menu p-4 overflow-y-auto w-80 bg-base-100 text-base-content">
      <div class="mb-5">
        <div class="flex items-center justify-between">
          <a href="#" class=" p-1.5 rounded flex items-center justify-center">
            <img src="https://tavuli.com/logo-green.svg" alt="">
          </a>
          <button class="flex items-center justify-center p-0.5 rounded bg-gray-800">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 stroke-current text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M7 16l-4-4m0 0l4-4m-4 4h18" />
            </svg>
          </button>
        </div>
      </div>
      <li>
        <a use:link href="/" class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Dashboard</a>
      </li>
      <li>
        <div on:click={toggletogel} class="relative flex justify-between text-gray-500 hover:bg-[#e6f7ff] hover:text-[#1890ff]">
          <div class="flex items-center w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
              </svg>
            </div>
            <a
              class="inline-block w-full py-2 pl-8 pr-4 text-sm rounded " 
              href="#">TOGEL</a>
          </div>
          <button class="absolute right-0 flex items-center p-1">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 stroke-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
        </div>
        {#if menu_togel_child_flag}
        <div class="pt-2 pl-4 hover:bg-base-100 cursor-none">
          <ul class="flex flex-col pl-2 text-gray-500 border-l border-[#e6f7ff] w-full">
            <li>
              <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">PERIODE</a>
            </li>
            <li>
              <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">PREDIKSI</a>
            </li>
            <li>
              <a href="#" class="inline-block w-full px-4 py-2 text-xs border-none rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">REPORT WINLOSE</a>
            </li>

          </ul>
        </div>
        {/if}
      </li>
      <li>
        <div on:click={togglemaster} class="relative flex justify-between text-gray-500 hover:bg-[#e6f7ff] hover:text-[#1890ff]">
          <div class="flex items-center w-full">
            <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
              </svg>
            </div>
            <a
              class="inline-block w-full py-2 pl-8 pr-4 text-sm rounded " 
              href="#">MASTER</a>
          </div>
          <button class="absolute right-0 flex items-center p-1">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 stroke-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
              <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
            </svg>
          </button>
        </div>
        {#if menu_master_child_flag}
        <div class="pt-2 pl-4 hover:bg-base-100 cursor-none">
          <ul class="flex flex-col pl-2 text-gray-500 border-l border-[#e6f7ff] w-full">
            <li>
              <a use:link href="/admin" class="inline-block w-full px-4 py-2 text-xs rounded border-none hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">PASARAN</a>
            </li>
            <li>
              <a use:link href="/admin" class="inline-block w-full px-4 py-2 text-xs rounded border-none hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">LOG MANAGEMENT</a>
            </li>
            <li>
              <a use:link href="/admin" class="inline-block w-full px-4 py-2 text-xs rounded border-none hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">ADMIN MANAGEMENT</a>
            </li>
            <li>
              <a use:link href="/adminrule" class="inline-block w-full px-4 py-2 text-xs rounded border-none active:bg-[#e6f7ff] active:outline-none hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-[#e6f7ff] focus:text-white">ADMIN RULE</a>
            </li>
          </ul>
        </div>
        {/if}
      </li>
      <li><a class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Sidebar</a></li>
    </ul>  
</div>