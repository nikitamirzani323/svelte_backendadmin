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