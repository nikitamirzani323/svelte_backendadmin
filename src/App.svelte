<script>
  let drawer_class = "drawer-mobile"
  drawer_class = ""
</script>

<main class="flex space-x-2 h-screen w-full bg-[#f0f2f5]">
  <div class="drawer {drawer_class}">
    <input id="my-drawer-2" type="checkbox" class="drawer-toggle">
    <div class="drawer-content flex flex-col gap-6">
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
            <button class="btn btn-ghost btn-circle">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white fill-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
              </svg>
            </button>
          </div>
        </div>
      </div>
      <div class="container mx-auto px-2">
        <div class="bg-white shadow-lg p-5 mt-5 mb-10">
          <div class="flex flex-col">
            <div class="flex items-start">
              <h1 class=" text-black font-bold text-[2.25em] lg:text-[2.25em] md:text-[2.25em] sm:text-[2.25em]">Admin</h1>
              <div class="flex flex-1 justify-end items-center gap-2 mt-2">
                <button class="hidden btn btn-primary rounded-md lg:inline-flex">New</button>
                <button class="hidden btn btn-primary rounded-md lg:inline-flex">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                  </svg>
                </button>
              </div>
            </div>
            <div class="mt-2 flex gap-2">
              <div class="relative w-full">
                <div class="absolute inset-y-0 left-0 flex items-center pl-2">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                    <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                  </svg>
                </div>
                <input type="text" placeholder="Search" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4">
              </div>
              <select class="select select-bordered  max-w-xs rounded-md">
                <option selected>PAGE 1 OF 50</option>
                <option>PAGE 2 OF 50</option>
                <option>PAGE 3 OF 50</option>
                <option>PAGE 4 OF 50</option>
                <option>PAGE 5 OF 50</option>
              </select>
            </div>
            <div class="mt-2 h-[650px]">
              <div class="overflow-x-auto">
                <table class="table table-compact w-full">
                  <thead>
                    <tr>
                      <th>&nbsp;</th>
                      <th class="text-center text-[11px] lg:text-sm">NO</th> 
                      <th>&nbsp;</th> 
                      <th class="text-center text-[11px] lg:text-sm">DATE</th>
                      <th class="text-left text-[11px] lg:text-sm">INVOICE</th> 
                      <th class="text-left text-[11px] lg:text-sm">PERIODE</th> 
                      <th class="text-left text-[11px] lg:text-sm">PASARAN</th> 
                      <th class="text-center text-[11px] lg:text-sm">PRIZE 1</th>
                      <th class="text-right text-[11px] lg:text-sm">REVISI</th>
                      <th class="text-right text-[11px] lg:text-sm">MEMBER</th>
                      <th class="text-right text-[11px] lg:text-sm">BET</th>
                      <th class="text-right text-[11px] lg:text-sm">BAYAR</th>
                      <th class="text-right text-[11px] lg:text-sm">CANCEL</th>
                      <th class="text-right text-[11px] lg:text-sm">WINLOSE</th>
                    </tr>
                  </thead> 
                  <tbody>
                    <tr>
                      <th>
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                          <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                        </svg>
                      </th> 
                      <th class="text-center  text-[11px] lg:text-sm">1</th> 
                      <td  class="text-center  text-[11px] lg:text-sm">
                        <div class="bg-yellow-400 text-center rounded-md p-2 font-bold">RUNNING</div>
                      </td> 
                      <td class="text-center  text-[11px] lg:text-sm">2022-04-03</td> 
                      <td class="text-left text-[11px] lg:text-sm">220402048</td> 
                      <td class="text-left text-[11px] lg:text-sm">73-DM</td> 
                      <td class="text-left text-[11px] lg:text-sm">DEMO</td> 
                      <td class="text-center text-[11px] lg:text-sm">8127</td>
                      <td class="text-right text-red-500 text-[11px] lg:text-sm">0</td>
                      <td class="text-right text-blue-500 text-[11px] lg:text-sm">1</td>
                      <td class="text-right text-blue-500 text-[11px] lg:text-sm">150</td>
                      <td class="text-right text-blue-500 text-[11px] lg:text-sm">536.050</td>
                      <td class="text-right text-red-500 text-[11px] lg:text-sm">0</td>
                      <td class="text-right text-blue-500 text-[11px] lg:text-sm">535.050</td>
                    </tr>
                   
                  </tbody> 
                </table>
              </div>
            </div>
            <div class="bg-gray-200 h-16 p-5">
              <span class="font-bold">TOTAL ROW : 100</span>
            </div>
          </div>
        </div>
      </div>
      
    </div> 
  
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
          <a class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Dashboard</a>
        </li>
        <li>
          <div class="relative flex justify-between text-gray-500 hover:bg-[#e6f7ff] hover:text-[#1890ff]">
            <div class="flex items-center w-full">
              <div class="absolute inset-y-0 left-0 flex items-center pl-2 pointer-events-none">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 stroke-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                  <path stroke-linecap="round" stroke-linejoin="round" d="M7 8h10M7 12h4m1 8l-4-4H5a2 2 0 01-2-2V6a2 2 0 012-2h14a2 2 0 012 2v8a2 2 0 01-2 2h-3l-4 4z" />
                </svg>
              </div>
              <a
                class="inline-block w-full py-2 pl-8 pr-4 text-xs rounded " 
                href="#">Content</a>
            </div>
            <button class="absolute right-0 flex items-center p-1">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 stroke-current" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M19 9l-7 7-7-7" />
              </svg>
            </button>
          </div>
          <div class="pt-2 pl-4 hover:bg-base-100 cursor-none">
            <ul class="flex flex-col pl-2 text-gray-500 border-l border-gray-700 w-full">
              <li>
                <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">Course</a>
              </li>
              <li>
                <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">Course</a>
              </li>
              <li>
                <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">Course</a>
              </li>
              <li>
                <a href="#" class="inline-block w-full px-4 py-2 text-xs rounded hover:bg-[#e6f7ff] hover:text-[#1890ff] focus:outline-none focus:ring-1 focus:ring-gray-500 focus:text-white">Course</a>
              </li>
            </ul>
          </div>
        </li>
        <li><a class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Sidebar Item 2</a></li>
        <li><a class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Sidebar Item 2</a></li>
        <li><a class="hover:bg-[#e6f7ff] hover:text-[#1890ff]">Sidebar Item 2</a></li>
      </ul>
    
    </div>
  </div>
  
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
