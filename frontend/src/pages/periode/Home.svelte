<script>
    import { createEventDispatcher } from "svelte";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_alert from '../../components/Modal_alert.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let totalrecord = 0;

    let page = "Periode";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Form_MemberlistBet = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modal_width_listbetmember = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";
    let searchMember = "";
    let searchMemberListBet = "";
    let searchHome = "";
    let filterHome = [];
    let filterMember = [];
    let filterMemberListBet = [];
    let listMember = [];
    let listBetUsername = [];

    let idtrxkeluaran = "";
    let idpasarancode = "";
    let periode_tglkeluaran_field = "";
    let periode_tanggalnext_field = "";
    let periode_periode_field = "";
    let periode_keluaran_field = "";
    let periode_status_field = "";
    let periode_statusonline_field = "";
    let periode_statusrevisi_field = "";
    let periode_create_field = "";
    let periode_createdate_field = "";
    let periode_update_field = "";
    let periode_updatedate_field = "";

    let totalbet = 0;
    let totalbayar = 0;
    let totalwin = 0;
    let total_member = 0;
    let subtotal_member_bet = 0;
    let subtotal_member_bayar = 0;
    let subtotal_member_cancel = 0;
    let subtotal_member_win = 0;
    let subtotal_member_winlose = 0;
    let subtotal_member_winlose_class = "text-blue-700";
    let client_username = ""

    let tab_listmember = "bg-sky-600 text-white"
    let tab_betgroup = ""
    let panel_listmember = true
    let panel_betgroup = false

    let dispatch = createEventDispatcher();
    
    async function SaveTransaksi() {
        let flag = false;
        msg_error = "";
        if (periode_keluaran_field == "") {
            flag = true;
            msg_error += "The Prize 1 is required\n";
        }
        if (parseInt(periode_keluaran_field.length) < 4) {
            flag = true;
            msg_error += "The Prize 1 is must 4 Character\n";
        }
        if (flag == false) {
            periode_status_field = "LOCK";
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/saveperiode", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page: "PERIODE-SAVE",
                    idinvoice: parseInt(idtrxkeluaran),
                    idpasarancode: idpasarancode,
                    nomorkeluaran: periode_keluaran_field,
                }),
            });
            const json = await res.json();
            if(!res.ok){
                loader_msg = "System Mengalami Trouble"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
            }else{
                if (json.status == 200) {
                    loader_msg = json.message
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                RefreshHalaman();
            }
        } else {
            alert(msg_error);
        }
    }
   
    async function EditData(e,y) {
        if(e != ""){
            isModalLoading = true;
            modal_width = "max-w-5xl";
            sData = "Edit";
            clearField();
            idtrxkeluaran = e;
            idpasarancode = y;
            const res = await fetch(path_api+"api/editperiode", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    idinvoice: parseInt(e),
                }),
            });
            const json = await res.json();
            let record = json.record;
            if (json.status === 400) {
                dispatch("handleLogout", "call");
            }else if(json.status === 200) {
                isModal_Form_New = true;
                isModalLoading = false;
                call_listmember()
                for (let i = 0; i < record.length; i++) {
                    let keluaran = record[i]["periode_keluaran"];
                    let document_status = "OPEN";
                    if (keluaran != "") {
                        document_status = "LOCK";
                    }
                    periode_status_field = document_status;
                    periode_tglkeluaran_field = record[i]["periode_tanggalkeluaran"];
                    periode_tanggalnext_field = record[i]["periode_tanggalnext"];
                    periode_periode_field = record[i]["periode_keluaranperiode"];
                    periode_keluaran_field = record[i]["periode_keluaran"];
                    periode_statusrevisi_field = record[i]["periode_statusrevisi"];
                    periode_statusonline_field = record[i]["periode_statusonline"];
                    periode_create_field = record[i]["periode_create"];
                    periode_createdate_field = record[i]["periode_createdate"];
                    periode_update_field = record[i]["periode_update"];
                    periode_updatedate_field = record[i]["periode_updatedate"];
                }
            }else{
                isModalLoading = false;
                isModalNotif = true;
                msg_error = "Silahkan Hubungi Administrator"
            }
        }
    }
    async function call_listmember() {
        listMember = [];
        const res = await fetch(path_api+"api/periodelistmember", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
            }),
        });
        const json = await res.json();
        let record = json.record;
        total_member = 0;
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            subtotal_member_bet = 0;
            subtotal_member_bayar = 0;
            subtotal_member_cancel = 0;
            subtotal_member_win = 0;
            subtotal_member_winlose = 0;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    total_member = total_member + 1;
                    subtotal_member_bet = subtotal_member_bet + parseInt(record[i]["totalbet"]);
                    subtotal_member_bayar = subtotal_member_bayar + parseInt(record[i]["totalbayar"]);
                    subtotal_member_cancel = subtotal_member_cancel + parseInt(record[i]["totalcancelbet"]);
                    subtotal_member_win = subtotal_member_win + parseInt(record[i]["totalwin"]);
                    listMember = [
                        ...listMember,
                        {
                            member_name: record[i]["member"],
                            member_totalbet: record[i]["totalbet"],
                            member_totalbayar: record[i]["totalbayar"],
                            member_totalcancel: record[i]["totalcancelbet"],
                            member_totalwin: record[i]["totalwin"],
                        },
                    ];
                }
                subtotal_member_winlose = subtotal_member_bayar - subtotal_member_cancel - subtotal_member_win;
                if (parseInt(subtotal_member_winlose) > 0) {
                    subtotal_member_winlose_class = "text-blue-700";
                }else{
                    subtotal_member_winlose_class = "text-red-500";
                }
            } else {
                setTimeout(function () {
                    let msgloader = "";
                    let css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    async function call_listbetbyusername(e) {
        listBetUsername = [];
        client_username = e.toUpperCase();
        modal_width_listbetmember = "max-w-7xl"
        isModal_Form_MemberlistBet = true;
        const res = await fetch(path_api+"api/periodelistbetusername", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                username: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            
        } else {
            if (record != null) {
                let status_class = ""
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["bet_status"] == "LOSE"){
                        status_class = "bg-red-400 text-white"
                    }else if(record[i]["bet_status"] == "CANCEL"){
                        status_class = "bg-red-600 text-white"
                    }else{
                        status_class = "bg-[#8BC34A] text-black"
                    }
                    listBetUsername = [
                        ...listBetUsername,
                        {
                            bet_id: record[i]["bet_id"],
                            bet_datetime: record[i]["bet_datetime"],
                            bet_ipaddress: record[i]["bet_ipaddress"],
                            bet_device: record[i]["bet_device"],
                            bet_timezone: record[i]["bet_timezone"],
                            bet_username: record[i]["bet_username"],
                            bet_typegame: record[i]["bet_typegame"],
                            bet_nomortogel: record[i]["bet_nomortogel"],
                            bet_posisitogel: record[i]["bet_posisitogel"],
                            bet_bet: record[i]["bet_bet"],
                            bet_diskon: record[i]["bet_diskon"],
                            bet_diskonpercen: record[i]["bet_diskonpercen"].toFixed(2),
                            bet_kei: record[i]["bet_kei"],
                            bet_keipercen: record[i]["bet_keipercen"].toFixed(2),
                            bet_bayar: record[i]["bet_bayar"],
                            bet_win: record[i]["bet_win"].toFixed(2),
                            bet_totalwin: record[i]["bet_totalwin"],
                            bet_status: record[i]["bet_status"],
                            bet_status_class: status_class,
                        },
                    ];
                }
            } else {
                setTimeout(function () {
                    let msgloader = "";
                    let css_loader = "display: none;";
                }, 1000);
            }
        }
    }
    const ChangeTabMenu = (e) => {
        switch(e){
            case "menu_listmember":
                tab_listmember = "bg-sky-600 text-white"
                tab_betgroup = ""
                panel_listmember = true
                panel_betgroup = false
                break;
            case "menu_betgroup":
                tab_listmember = ""
                tab_betgroup = "bg-sky-600 text-white"
                panel_listmember = false
                panel_betgroup = true
                break;
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
   
    function clearField(){
        idtrxkeluaran = "";
        idpasarancode = "";
        periode_tglkeluaran_field = "";
        periode_tanggalnext_field = "";
        periode_periode_field = "";
        periode_keluaran_field = "";
        periode_status_field = "";
        periode_statusonline_field = "";
        periode_statusrevisi_field = "";
        periode_create_field = "";
        periode_createdate_field = "";
        periode_update_field = "";
        periode_updatedate_field = "";
    }
    
    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_invoice
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_status
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
        if (searchMember) {
            filterMember = listMember.filter(
                (item) =>
                    item.member_name
                        .toLowerCase()
                        .includes(searchMember.toLowerCase()) 
            );
        } else {
            filterMember = [...listMember];
        }
        if (searchMemberListBet) {
            filterMemberListBet = listBetUsername.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchMemberListBet.toLowerCase()) || 
                    item.bet_typegame
                        .toLowerCase()
                        .includes(searchMemberListBet.toLowerCase())
            );
        } else {
            filterMemberListBet = [...listBetUsername];
        }
    }
</script>
<div class="{loader_class} w-40 fixed top-1 left-[45%] right-0 z-50">
    <div class="alert alert-warning shadow-lg rounded-md bg-[#FFF6BF]">
        <div>
          <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
          <span>{loader_msg}</span>
        </div>
    </div>
</div>
<div class="container mx-auto px-2 lg:px-28">
    <div class="bg-white shadow-lg p-5">
        <div class="flex flex-col gap-2">
            <div class="flex items-start">
                <h1 class=" text-slate-600 font-bold text-sm lg:text-3xl uppercase w-full">{page}</h1>
                <div class="hidden sm:flex md:flex justify-end w-full gap-2 ">
                    <button on:click={() => {
                        RefreshHalaman();
                        }} class="btn btn-primary hover:bg-primary shadow-lg shadow-[#b3e4fc]  rounded-md lg:inline-flex">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                            <path stroke-linecap="round" stroke-linejoin="round" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
                        </svg>
                    </button>
                </div>
            </div>
            <div class="relative w-full">
                <div class="absolute inset-y-0 left-0 flex items-center pl-2">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                        <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                    </svg>
                </div>
                <input 
                    bind:value={searchHome}
                    type="text" placeholder="Search by Rule" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 ">
            </div>
            <div class="hidden sm:inline w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[550px] overflow-y-scroll">
                <table class="table table-compact w-full ">
                    <thead class="sticky top-0">
                        <tr>
                            <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                            <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NO</th>
                            <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center"></th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">DATE</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">INVOICE</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PERIODE</th>
                            <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PASARAN</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">PRIZE 1</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">REVISI</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">MEMBER</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">BET</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">BAYAR</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">CANCEL</th>
                            <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WINLOSE</th>
                        </tr>
                    </thead>
                    {#if filterHome != ""}
                        <tbody>
                            {#each filterHome as rec}
                            <tr>
                                <td on:click={() => {
                                    EditData(rec.home_invoice,rec.home_code);
                                    }} class="text-center text-xs cursor-pointer">
                                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                        <path stroke-linecap="round" stroke-linejoin="round" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                                    </svg>
                                </td>
                                <td class="text-xs lg:text-sm align-top text-center">{rec.home_no}</td>
                                <td class="text-xs lg:text-sm align-top text-center">
                                    <span class="{rec.home_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.home_status}</span>
                                </td>
                                <td class="text-xs lg:text-sm align-top text-center">{rec.home_tanggal}</td>
                                <td class="text-xs lg:text-sm align-top text-left">{rec.home_invoice}</td>
                                <td class="text-xs lg:text-sm align-top text-left">{rec.home_periode}</td>
                                <td class="text-xs lg:text-sm align-top text-left">{rec.home_name}</td>
                                <td class="text-xs lg:text-sm align-top text-center font-semibold">{rec.home_keluaran}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_revisi_class}">{new Intl.NumberFormat().format(rec.home_revisi)}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_totalmember_class}">{new Intl.NumberFormat().format(rec.home_totalmember)}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_totalbet_class}">{new Intl.NumberFormat().format(rec.home_totalbet)}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_totaloutstanding_class}">{new Intl.NumberFormat().format(rec.home_totaloutstanding)}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_totalcancelbet_class}">{new Intl.NumberFormat().format(rec.home_totalcancelbet)}</td>
                                <td class="text-xs lg:text-sm align-top text-right font-semibold {rec.home_winlose_class}">{new Intl.NumberFormat().format(rec.home_winlose)}</td>
                            </tr>
                            {/each}
                        </tbody>
                    {:else}
                        <tbody>
                            <tr>
                                <td colspan="14" class="text-center">
                                    <progress class="self-start progress progress-primary w-56"></progress>
                                </td>
                            </tr>
                        </tbody>
                    {/if}
                </table>
            
            </div>
            
            <div class="bg-[#F7F7F7] rounded-sm h-16 p-5">
                <span class="font-bold">TOTAL ROW : {totalrecord}</span>
            </div>
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<div class="modal" >
    <div class="modal-box relative select-none w-11/12 {modal_width}  rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-formnew" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">Entry/{sData}</h3>
            </div>
            {#if sData=="New"}
                <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
                    <div class="relative form-control mt-2">
                        
                    </div>
                </div>
                <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                    <button class="{buttonLoading_class}">Submit</button>
                </div>
            {/if}
            {#if sData=="Edit"}
                <div class="flex justify-between  gap-2">
                    <div class="w-1/2">
                        <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2  ">
                            <div class="alert alert-info shadow-lg mt-2 rounded-sm">
                                <div>
                                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                                  <span>Periode Selanjutnya : {periode_tanggalnext_field}</span>
                                </div>
                            </div>
                            <div class="relative form-control">
                                <Input_custom
                                    input_enabled={false}
                                    input_tipe="text"
                                    bind:value={idtrxkeluaran}
                                    input_id="idtrxkeluaran"
                                    input_placeholder="Invoice"/>
                            </div>
                            <div class="relative form-control">
                                <Input_custom
                                    input_enabled={false}
                                    input_tipe="tanggal"
                                    bind:value={periode_tglkeluaran_field}
                                    input_id="periode_tglkeluaran_field"
                                    input_placeholder="Tanggal"/>
                            </div>
                            <div class="relative form-control">
                                <Input_custom
                                    input_enabled={false}
                                    input_tipe="text"
                                    bind:value={periode_periode_field}
                                    input_id="periode_periode_field"
                                    input_placeholder="Periode"/>
                            </div>
                            <div class="relative form-control">
                                <Input_custom
                                    input_required={true}
                                    input_tipe="number_nolabel"
                                    bind:value={periode_keluaran_field}
                                    input_maxlenght="4"
                                    input_id="periode_keluaran_field"
                                    input_placeholder="Prize 1"/>
                            </div>
                            <div class="relative form-control text-[11px]">
                                Create : {periode_create_field}, {periode_createdate_field}
                                {#if periode_update_field != ""}
                                    <br>
                                    Update : {periode_update_field}, {periode_updatedate_field}
                                {/if}
                            </div>
                        </div>
                        {#if periode_status_field == "OPEN"}
                            {#if periode_statusonline_field == "OFFLINE"}
                                <div class="flex flex-wrap justify-end align-middle  mt-2">
                                    <button
                                        on:click={() => {
                                            SaveTransaksi();
                                        }}  
                                        class="{buttonLoading_class} btn-block">Submit</button>
                                </div>
                            {/if}
                        {/if}
                    </div>
                    <div class="w-full p-2">
                        <ul class="flex justify-center items-center gap-2">
                            <li on:click={() => {
                                    ChangeTabMenu("menu_listmember");
                                }}
                                class="items-center {tab_listmember}  px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Member</li>
                            <li on:click={() => {
                                    ChangeTabMenu("menu_betgroup");
                                }}
                                class="items-center {tab_betgroup} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">Bet Group</li>
                        </ul>
                        {#if panel_listmember}
                            <h2 class="text-lg font-bold mb-2">List Member</h2>
                            <input 
                                bind:value={searchMember}
                                type="text" placeholder="Search by Username" class="input input-bordered w-full max-w-full rounded-md  ">
                            <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[400px] overflow-y-scroll mt-2">
                                <table class="table table-compact w-full">
                                    <thead class="sticky top-0">
                                        <tr>
                                            <th class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                            <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL<br>BET</th>
                                            <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL<br>BAYAR</th>
                                            <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL<br>CANCEL</th>
                                            <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL<br>WIN</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        {#if filterMember != ""}
                                            {#each filterMember as rec}
                                                <tr>
                                                    <td on:click={() => {
                                                        call_listbetbyusername(rec.member_name);
                                                    }} class="text-xs text-left align-top underline cursor-pointer">{rec.member_name}</td>
                                                    <td class="text-xs text-right align-top text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.member_totalbet)}</td>
                                                    <td class="text-xs text-right align-top text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.member_totalbayar)}</td>
                                                    <td class="text-xs text-right align-top text-red-500 font-semibold">{new Intl.NumberFormat().format(rec.member_totalcancel)}</td>
                                                    <td class="text-xs text-right align-top text-red-500 font-semibold">{new Intl.NumberFormat().format(rec.member_totalwin)}</td>
                                                </tr>
                                            {/each}
                                        {:else}
                                            <tr>
                                                <td colspan="5" class="text-xs">No Records</td>
                                            </tr>
                                        {/if}
                                    </tbody>
                                </table>
                            </div>
                            <div class="">
                                <table class=" w-full">
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL MEMBER</td>
                                        <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(total_member)}</td>
                                    </tr>
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL BET</td>
                                        <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_member_bet)}</td>
                                    </tr>
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL BAYAR</td>
                                        <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_member_bayar)}</td>
                                    </tr>
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL CANCEL</td>
                                        <td class="text-xs font-semibold text-right align-top text-blue-700">{new Intl.NumberFormat().format(subtotal_member_cancel)}</td>
                                    </tr>
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL WIN</td>
                                        <td class="text-xs font-semibold text-right align-top text-red-500">{new Intl.NumberFormat().format(subtotal_member_win)}</td>
                                    </tr>
                                    <tr>
                                        <td class="text-xs font-semibold text-left align-top">TOTAL WINLOSE</td>
                                        <td class="text-xs font-semibold text-right align-top {subtotal_member_winlose_class}">{new Intl.NumberFormat().format(subtotal_member_winlose)}</td>
                                    </tr>
                                </table>
                            </div>
                        {/if}
                    </div>
                </div>
            {/if}
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-formlistbetmember" class="modal-toggle" bind:checked={isModal_Form_MemberlistBet}>
<div class="modal" >
    <div class="modal-box relative  w-11/12 {modal_width_listbetmember}  rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-formlistbetmember" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">INFORMATION : {client_username}</h3>
            </div>
            <input 
                bind:value={searchMemberListBet}
                type="text" placeholder="Search by Status,Code" class="input input-bordered mt-1 w-full max-w-full rounded-md  ">
            <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[550px] overflow-y-scroll mt-2">
                <table class="table table-compact w-full">
                    <thead class="sticky top-0">
                        <tr>
                            <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                            <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">CODE</th>
                            <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">TANGGAL</th>
                            <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                            <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">IPADDRESS</th>
                            <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">DEVICE</th>
                            <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIPE</th>
                            <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">PERMAINAN</th>
                            <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">NOMOR</th>
                            <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BET</th>
                            <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">DISC</th>
                            <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">KEI</th>
                            <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BAYAR</th>
                            <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN</th>
                            <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN<br>TOTAL</th>
                        </tr>
                    </thead>
                    <tbody>
                        {#if filterMemberListBet != ""}
                            {#each filterMemberListBet as rec}
                                <tr>
                                    <td class="text-xs text-center align-top">
                                        <span class="{rec.bet_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.bet_status}</span>  
                                    </td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_id}</td>
                                    <td class="text-xs text-center align-top whitespace-nowrap">{rec.bet_datetime}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_username}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_ipaddress}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_device}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_posisitogel}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_typegame}</td>
                                    <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_nomortogel}</td>
                                    <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                                    <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_diskon)}</td>
                                    <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_kei)}</td>
                                    <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                    <td class="text-xs text-right align-top font-semibold whitespace-nowrap">{rec.bet_win}x</td>
                                    <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
                                </tr>
                            {/each}
                        {:else}
                            <tr>
                                <td colspan="15" class="text-xs text-center">
                                    <progress class="self-start progress progress-primary w-56"></progress>
                                </td>
                            </tr>
                        {/if}
                    </tbody>
                </table>
            </div>
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-notif" class="modal-toggle" bind:checked={isModalNotif}>
<Modal_alert 
    modal_tipe="notifikasi" modal_id="my-modal-notif" 
    modal_title="Information" modal_message="{msg_error}"  />

<input type="checkbox" id="my-modal-loading" class="modal-toggle" bind:checked={isModalLoading}>
<Modal_alert modal_tipe="loading" modal_widthheight_class="w-auto grass opacity-50"  />
