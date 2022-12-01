<script>
    import { createEventDispatcher } from "svelte";
    import Loader from '../../../components/Loader.svelte' 
    import Input_custom from '../../../components/Input.svelte' 
    import Panel_info from '../../../components/Panel_info.svelte' 
    import Panel_table from '../../../components/panel_table.svelte' 
    import Modal_popup from '../../../components/Modal_popup.svelte'
    export let path_api = "";
    export let token = "";
    export let font_size = "";
    export let listpasaran = [];
    let sData = "Edit";
    let isModalNotif = false
    let isModalLoading = false
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let msg_error = "";
    let buttonLoading_class = "btn btn-primary"
    let isModal_Form_Submit = false
    let isModal_Form_generator = false
    let isModal_Form_listBetmember = false
    let modal_width = "max-w-xl"
    let modal_width_listbetmember = "max-w-xl"

    let label_invoice = "";

    let idtrxkeluaran = "";
    let idpasarancode = "";
    let pasaran_msgrevisi = "";
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

    let generator_member_field = 0;
    let generator_bet_field = 0;

    let listPeriodeMember = [];
    let searchListMember = "";
    let filterListBetMember = [];
    let total_member = 0;
    let subtotal_member_bet = 0;
    let subtotal_member_bayar = 0;
    let subtotal_member_cancel = 0;
    let subtotal_member_win = 0;
    let subtotal_member_winlose = 0;
    let subtotal_member_winlose_class = "";

    let dispatch = createEventDispatcher();
    const handleAction = (type,idincoive,idpasaran) => {
        idtrxkeluaran = "";
        switch(type){
            case "LIVEDRAW":
                window.open(idincoive,'_blank');
                break;
            case "SUBMIT":
                EditDataPeriode(idincoive)
                modal_width = "max-w-xl";
                isModal_Form_Submit = true;
                idpasarancode = idpasaran
                break;
            case "GENERATOR":
                label_invoice = idincoive
                idtrxkeluaran = idincoive
                modal_width = "max-w-xl";
                isModal_Form_generator = true;
                generator_member_field = 0;
                generator_bet_field = 0;
                break;
            case "LISTMEMBER":
                call_periodelistmember(idincoive)
                modal_width_listbetmember = "max-w-2xl";
                isModal_Form_listBetmember = true;
                break;
        }
    };
    const cancelbetTransaksi = (type,e) => {
        
    };
    async function EditDataPeriode(e) {
        if(e != ""){
            label_invoice = e
            isModalLoading = true;
            modal_width = "max-w-5xl";
            sData = "Edit";
            clearFieldPeriode();
            idtrxkeluaran = e;
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
                // dispatch("handleLogout", "call");
            }else if(json.status === 200) {
                isModal_Form_Submit = true;
                isModalLoading = false;
               
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
    async function call_periodelistmember(e) {
        resetlistmember()
        const res = await fetch(path_api+"api/periodelistmember", {
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
        total_member = 0;
        if (json.status === 400) {
            // dispatch("handleLogout", "call");
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
                    listPeriodeMember = [
                        ...listPeriodeMember,
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
                isModal_Form_Submit = false;
                RefreshHalamanPasaran();
                // EditData(idtrxkeluaran,idpasarancode,pasaran_msgrevisi)
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true
            }
        }
    }
    async function SaveGenerator() {
        let flag = false;
        msg_error = "";
        if (generator_member_field < 0) {
            flag = true;
            msg_error += "The Total Member is required\n";
        }
        if (generator_bet_field < 0) {
            flag = true;
            msg_error += "The Total Bet is required\n";
        }
        if (flag == false) {
            periode_status_field = "LOCK";
            buttonLoading_class = "btn loading"
            loader_class = "inline-block"
            loader_msg = "Sending..."
            const res = await fetch(path_api+"api/generatorsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page: "PERIODE-SAVE",
                    invoice: idtrxkeluaran.toString(),
                    totalmember: parseInt(generator_member_field),
                    totalrow: parseInt(generator_bet_field),
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
                    loader_msg = "Harap Tunggu Beberapa Menit"
                } else if (json.status == 403) {
                    loader_msg = json.message
                } else {
                    loader_msg = json.message;
                }
                buttonLoading_class = "btn btn-primary"
                setTimeout(function () {
                    loader_class = "hidden";
                }, 1000);
                isModal_Form_generator = false;
                // RefreshHalamanPasaran();
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true
            }
        }
    }
    function clearFieldPeriode(){
        idtrxkeluaran = "";
        idpasarancode = "";
        pasaran_msgrevisi = "";
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
    function resetlistmember(){
        listPeriodeMember = [];
        searchListMember = "";
        filterListBetMember = [];
        total_member = 0;
        subtotal_member_bet = 0;
        subtotal_member_bayar = 0;
        subtotal_member_cancel = 0;
        subtotal_member_win = 0;
        subtotal_member_winlose = 0;
        subtotal_member_winlose_class = "";
    }
    const RefreshHalamanPasaran = () => {
        dispatch("handleRefreshDataListPasaran", "call");
    };
    $: {
        if (searchListMember) {
            filterListBetMember = listPeriodeMember.filter(
                (item) =>
                    item.member_name
                        .toLowerCase()
                        .includes(searchListMember.toLowerCase()) 
            );
        } else {
            filterListBetMember = [...listPeriodeMember];
        }
    }
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
<div class="mt-1 gap-2 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4 mx-2  ">
    {#each listpasaran as rec}
        <div class="card  bg-base-100 shadow-xl rounded-md  z-0">
            <div class="card-body p-0">
                <div class="card-title bg-primary p-2 select-none">
                    <div class="flex justify-items-center w-full gap-2">
                        <div class="w-full">
                            <h2 class="text-sm lg:text-lg font-bold text-white">{rec.home_nmpasaran}</h2>
                        </div>
                        <div class="dropdown dropdown-end hidden lg:grid">
                            <svg tabindex="0" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6 text-white cursor-pointer">
                                <path stroke-linecap="round" stroke-linejoin="round" d="M12 6.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 12.75a.75.75 0 110-1.5.75.75 0 010 1.5zM12 18.75a.75.75 0 110-1.5.75.75 0 010 1.5z" />
                            </svg>
                            <ul tabindex="0" class="dropdown-content menu p-2 shadow bg-base-100 rounded-md w-52 ">
                              <li><span class="text-xs lg:text-sm">NEW INVOICE</span></li>
                              <li><span class="text-xs lg:text-sm">CHART</span></li>
                              <li><span class="text-xs lg:text-sm">SETTING</span></li>
                              <li><span on:click={() => { handleAction("LIVEDRAW",rec.home_pasaranurl);}} class="text-xs lg:text-sm">LIVEDRAW</span></li>
                            </ul>
                        </div>
                    </div>
                </div>
                <div class="hidden lg:grid">
                    <div class="w-full p-2">
                        <div class="alert alert-info shadow-sm rounded-sm -my-2 w-full p-0 select-none">
                            <div class="flex flex-col w-full gap-0">
                                <center class="text-xs lg:text-lg text-black font-bold text-center w-full border-b-2 border-gray-200 py-2">
                                    {rec.home_diundi}
                                </center>
                                <div class="grid grid-cols-3 w-full p-2 ">
                                    <span class="text-xs lg:text-sm text-black font-bold text-center w-full">JAM JADWAL<br />{rec.home_jadwal}</span>
                                    <span class="text-xs lg:text-sm text-black font-bold text-center w-full">JAM OPEN<br />{rec.home_jadwal}</span>
                                    <span class="text-xs lg:text-sm text-black font-bold text-center w-full">JAM TUTUP<br />{rec.home_jadwal}</span>
                                  </div>
                            </div>
                        </div>
                        <table class="w-full mt-4">
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">DATE</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm font-semibold">{rec.home_periode_tglperiode}</td>
                            </tr>
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">INVOICE</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm font-semibold">{rec.home_periode_idinvoice}</td>
                            </tr>
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">PERIODE</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm font-semibold">{rec.home_periode_nomorperiode}</td>
                            </tr>
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">MEMBER</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.home_periode_total_member)}</td>
                            </tr>
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">TOTAL BET</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.home_periode_total_bet)}</td>
                            </tr>
                            <tr>
                                <td class="text-xs lg:text-sm font-bold">TOTAL BAYAR</td>
                                <td class="text-xs lg:text-sm font-bold">:</td>
                                <td class="text-xs lg:text-sm text-right text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.home_periode_total_bayar)}</td>
                            </tr>
                        </table>
                    </div>
                    <div class="gap-1 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-5 p-2 bg-base-200">
                        <button on:click={() => { 
                            handleAction("SUBMIT",rec.home_periode_idinvoice,rec.home_codepasarantogel);
                            }} class="btn btn-sm btn-primary text-xs rounded-sm p-1">Submit</button>
                        <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Simulasi</button>
                        <button on:click={() => { 
                            handleAction("LISTMEMBER",rec.home_periode_idinvoice,"");
                            }} class="btn btn-sm btn-primary text-xs rounded-sm p-1">List Member</button>
                        <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Bet Group</button>
                        <button class="btn btn-sm btn-primary text-xs rounded-sm p-1">Past Result</button>
                        <button on:click={() => { 
                            handleAction("GENERATOR",rec.home_periode_idinvoice,"");
                            }} class="btn btn-sm btn-primary text-xs rounded-sm p-1">Generator</button>
                    </div>
                </div>
            </div>
        </div>
    {/each}
</div>

<input type="checkbox" id="my-modal-formsubmit" class="modal-toggle" bind:checked={isModal_Form_Submit}>
<Modal_popup
    modal_popup_id="my-modal-formsubmit"
    modal_popup_title="INVOICE : {label_invoice}"
    modal_popup_class="select-none w-full lg:w-9/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData=="Edit"}
            <div class="w-full">
                <div class="flex flex-auto flex-col overflow-auto gap-2 mt-2  ">
                    {#if periode_status_field == "OPEN"}
                        <div class="alert alert-info shadow-lg mt-2 rounded-sm">
                            <div>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                            <span>Periode Selanjutnya : {periode_tanggalnext_field}</span>
                            </div>
                        </div>
                    {/if}
                    <div class="relative form-control mt-2">
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
                            input_tipe="number_nolabel_string"
                            bind:value={periode_keluaran_field}
                            input_maxlenght="4"
                            input_id="periode_keluaran_field"
                            input_placeholder="Prize 1"/>
                    </div>
                    <Panel_info>
                        <slot:template slot="panel_body">
                            <table>
                                <tr>
                                    <td>Create</td>
                                    <td>:</td>
                                    <td>{periode_create_field}, {periode_createdate_field}</td>
                                </tr>
                                {#if periode_update_field != ""}
                                <tr>
                                    <td>Modified</td>
                                    <td>:</td>
                                    <td>{periode_update_field}, {periode_updatedate_field}</td>
                                </tr>
                                {/if}
                            </table>
                        </slot:template>
                    </Panel_info>
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
        {/if}
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formgenerator" class="modal-toggle" bind:checked={isModal_Form_generator}>
<Modal_popup
    modal_popup_id="my-modal-formgenerator"
    modal_popup_title="INVOICE : {label_invoice} - GENERATOR"
    modal_popup_class="select-none w-full lg:w-9/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="w-full">
            <div class="flex flex-auto flex-col overflow-auto gap-2 mt-2  ">
                <div class="relative form-control mt-2">
                    <Input_custom
                        input_enabled={false}
                        input_tipe="text"
                        bind:value={idtrxkeluaran}
                        input_id="idtrxkeluaran"
                        input_placeholder="Invoice"/>
                </div>
                <div class="relative form-control">
                    <Input_custom
                        input_required={true}
                        input_tipe="number_nolabel_string"
                        input_maxlenght="4"
                        bind:value={generator_member_field}
                        input_id="generator_member_field"
                        input_placeholder="Total Member"/>
                </div>
                <div class="relative form-control">
                    <Input_custom
                        input_required={true}
                        input_tipe="number_nolabel_string"
                        input_maxlenght="4"
                        bind:value={generator_bet_field}
                        input_id="generator_bet_field"
                        input_placeholder="Total Bet"/>
                </div>
            </div>
            <div class="flex flex-wrap justify-end align-middle  mt-2">
                <button
                    on:click={() => {
                        SaveGenerator();
                    }}  
                    class="{buttonLoading_class} btn-block">Submit</button>
            </div>
        </div>
    </slot:template>
</Modal_popup>


<input type="checkbox" id="my-modal-listbetall" class="modal-toggle" bind:checked={isModal_Form_listBetmember}>
<div class="modal" >
    <div class="modal-box relative  w-11/12 {modal_width_listbetmember}  rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-listbetall" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">LIST BET</h3>
            </div>
            
            <Panel_table
                panel_class="mt-2" 
                panel_height="500px">
                <slot:template slot="paneltable_body">
                    <div class="flex justify-start items-stretch gap-2 mb-2 w-full">
                            
                        <div class="p-0 w-full">
                            <input type="text" placeholder="Search by Status, Code, Nomor" class="input input-bordered input-sm  rounded-sm w-full focus:ring-0 focus:outline-none">
                        </div>
                        <div class="p-0 w-32">
                            <select name="" id="" class="p-1 bg-gray-50 border border-gray-300 cursor-pointer">
                                <option value="">1 of 1000</option>
                            </select>
                        </div>
                    </div>
                    <table class="table table-compact w-full">
                        <thead class="sticky top-0">
                            <tr>
                                <th width="*" class="bg-[#475289] {font_size} text-white text-left align-top">USERNAME</th>
                                <th width="10%" class="bg-[#475289] {font_size} text-white text-right align-top">TOTAL BET</th>
                                <th width="10%" class="bg-[#475289] {font_size} text-white text-right align-top">TOTAL BAYAR</th>
                                <th width="10%" class="bg-[#475289] {font_size} text-white text-right align-top">TOTAL CANCEL</th>
                                <th width="10%" class="bg-[#475289] {font_size} text-white text-right align-top">TOTAL WIN</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#if filterListBetMember  != ""}
                                {#each filterListBetMember as rec}
                                    <tr>
                                        <td class="{font_size} text-left align-top whitespace-nowrap">{rec.member_name}</td>
                                        <td class="{font_size} text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_totalbet)}</td>
                                        <td class="{font_size} text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_totalbayar)}</td>
                                        <td class="{font_size} text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_totalcancel)}</td>
                                        <td class="{font_size} text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_totalwin)}</td>
                                    </tr>
                                {/each}
                            {:else}
                                <tr>
                                    <td colspan="16" class="text-xs text-center">
                                        <progress class="self-start progress progress-primary w-56"></progress>
                                    </td>
                                </tr>
                            {/if}
                        </tbody>
                    </table>
                    
                  
                </slot:template>
            </Panel_table>
            
            <div class="bg-[#F7F7F7] rounded-sm h-32 p-2">
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
        </div>
    </div>
</div>