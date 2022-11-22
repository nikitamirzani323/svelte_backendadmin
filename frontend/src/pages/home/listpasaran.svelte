<script>
    export let font_size = "";
    export let listpasaran = [];
</script>
<div class="mt-1 gap-1 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4 mx-2 ">
    {#each listpasaran as rec}
        <div class="card  bg-base-100 shadow-xl rounded-md select-none">
            <div class="card-body p-0">
                <div class="card-title bg-primary p-2">
                    <div class="flex justify-items-center w-full gap-2">
                        <div class="w-full">
                            <h2 class="text-sm lg:text-lg font-bold text-white">{rec.nmpasaran}</h2>
                        </div>
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white cursor-pointer">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M12 4.5v15m7.5-7.5h-15" />
                          </svg>
                          
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white cursor-pointer">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M3 13.125C3 12.504 3.504 12 4.125 12h2.25c.621 0 1.125.504 1.125 1.125v6.75C7.5 20.496 6.996 21 6.375 21h-2.25A1.125 1.125 0 013 19.875v-6.75zM9.75 8.625c0-.621.504-1.125 1.125-1.125h2.25c.621 0 1.125.504 1.125 1.125v11.25c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V8.625zM16.5 4.125c0-.621.504-1.125 1.125-1.125h2.25C20.496 3 21 3.504 21 4.125v15.75c0 .621-.504 1.125-1.125 1.125h-2.25a1.125 1.125 0 01-1.125-1.125V4.125z" />
                          </svg>
                          
                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white cursor-pointer">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M9.594 3.94c.09-.542.56-.94 1.11-.94h2.593c.55 0 1.02.398 1.11.94l.213 1.281c.063.374.313.686.645.87.074.04.147.083.22.127.324.196.72.257 1.075.124l1.217-.456a1.125 1.125 0 011.37.49l1.296 2.247a1.125 1.125 0 01-.26 1.431l-1.003.827c-.293.24-.438.613-.431.992a6.759 6.759 0 010 .255c-.007.378.138.75.43.99l1.005.828c.424.35.534.954.26 1.43l-1.298 2.247a1.125 1.125 0 01-1.369.491l-1.217-.456c-.355-.133-.75-.072-1.076.124a6.57 6.57 0 01-.22.128c-.331.183-.581.495-.644.869l-.213 1.28c-.09.543-.56.941-1.11.941h-2.594c-.55 0-1.02-.398-1.11-.94l-.213-1.281c-.062-.374-.312-.686-.644-.87a6.52 6.52 0 01-.22-.127c-.325-.196-.72-.257-1.076-.124l-1.217.456a1.125 1.125 0 01-1.369-.49l-1.297-2.247a1.125 1.125 0 01.26-1.431l1.004-.827c.292-.24.437-.613.43-.992a6.932 6.932 0 010-.255c.007-.378-.138-.75-.43-.99l-1.004-.828a1.125 1.125 0 01-.26-1.43l1.297-2.247a1.125 1.125 0 011.37-.491l1.216.456c.356.133.751.072 1.076-.124.072-.044.146-.087.22-.128.332-.183.582-.495.644-.869l.214-1.281z" />
                            <path stroke-linecap="round" stroke-linejoin="round" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                        </svg>
                    </div>
                </div>
                <div class="w-full p-2">
                    <table class="w-full">
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">DATE</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm font-semibold">2022-11-19</td>
                        </tr>
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">INVOICE</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm font-semibold">2211181588</td>
                        </tr>
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">PERIODE</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm font-semibold">315-DM</td>
                        </tr>
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">MEMBER</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">10</td>
                        </tr>
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">TOTAL BET</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">50.000</td>
                        </tr>
                        <tr>
                            <td class="text-xs lg:text-sm font-bold">TOTAL BAYAR</td>
                            <td class="text-xs lg:text-sm font-bold">:</td>
                            <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">10.000.000.000</td>
                        </tr>
                    </table>
                </div>
                <div class="gap-1 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-5 p-2 bg-base-200">
                    <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Submit</button>
                    <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Simulasi</button>
                    <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">List Member</button>
                    <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Bet Group</button>
                    <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Past Result</button>
                </div>
            </div>
        </div>
    {/each}
</div>