<script>
    import Loader from '../../components/Loader.svelte' 
    import Input_custom from '../../components/Input.svelte' 
    export let path_api = ""
    let listHome = [];
    let listresult  = [];
    let record = "";
    let prediksi_field = "";
    let prediksi_select_field = "";
    let token = localStorage.getItem("token");
    let akses_page = false;
    let isModalNotif = false;
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let totalbet = 0;
    let totalwin = 0;
    let subtotal = 0;
    let temp_companywinlose_class = "";
    let msg_error = ""
    async function initapp() {
        const res = await fetch(path_api+"api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "PREDIKSI-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            akses_page = true;
            initHome();
        }
    }
    async function initHome() {
        const res = await fetch(path_api+"api/listpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listHome = [
                        ...listHome,
                        {
                            home_idcomp: record[i]["pasaran_idcomp"],
                            home_code: record[i]["pasaran_code"],
                            home_nama: record[i]["pasaran_name"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function callPrediksi() {
        listresult = [];
        totalbet = 0;
        totalwin = 0;
        subtotal = 0;
        temp_companywinlose_class = "";
        let flag = false;
        msg_error = "";
        buttonLoading_class = "btn loading"
        loader_class = "inline-block"
        loader_msg = "Sending..."
        if (prediksi_select_field == "") {
            flag = true;
            msg_error += "The Pasaran is required\n";
        }
        if (prediksi_field == "") {
            flag = true;
            msg_error += "The Nomor Keluaran is required\n";
        }
        if (parseInt(prediksi_field.length) < 4) {
            flag = true;
            msg_error += "The Nomor Keluaran is must 4 Character\n";
        }
        if (flag == false) {
            const res = await fetch(path_api+"api/listprediksi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    nomorkeluaran: prediksi_field,
                    pasaran_code: parseInt(prediksi_select_field),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message;
                let record = json.record;
                totalbet = json.totalbet;
                totalwin = json.subtotal;
                subtotal = json.subtotalwin;
                if (parseInt(subtotal) > 0) {
                    temp_companywinlose_class = "text-blue-700 font-semibold";
                } else {
                    temp_companywinlose_class = "text-red-500 font-semibold";
                }
                if (record != null) {
                    let status_class = "";
                    for (var i = 0; i < record.length; i++) {
                        if(record[i]["prediksi_status"] == "LOSE"){
                            status_class = "bg-red-400 text-white"
                        }else if(record[i]["prediksi_status"] == "CANCEL"){
                            status_class = "bg-red-600 text-white"
                        }else if(record[i]["prediksi_status"] == "WINNER"){
                            status_class = "bg-[#8BC34A] text-black"
                        }else{
                            status_class = "bg-[#prediksi_status] text-black"
                        }
                        listresult = [
                            ...listresult,
                            {
                                prediksi_invoice: record[i]["prediksi_invoice"],
                                prediksi_code: record[i]["prediksi_code"],
                                prediksi_tanggal: record[i]["prediksi_tanggal"],
                                prediksi_username: record[i]["prediksi_username"],
                                prediksi_permainan: record[i]["prediksi_permainan"],
                                prediksi_posisitogel: record[i]["prediksi_posisitogel"],
                                prediksi_nomor: record[i]["prediksi_nomor"],
                                prediksi_bet: record[i]["prediksi_bet"],
                                prediksi_diskon: record[i]["prediksi_diskon"],
                                prediksi_diskonpercen: record[i]["prediksi_diskonpercen"].toFixed(2),
                                prediksi_kei: record[i]["prediksi_kei"],
                                prediksi_keipercen: record[i]["prediksi_keipercen"].toFixed(2),
                                prediksi_bayar: record[i]["prediksi_bayar"],
                                prediksi_win: record[i]["prediksi_win"].toFixed(2),
                                prediksi_totalwin: record[i]["prediksi_totalwin"],
                                prediksi_status: record[i]["prediksi_status"],
                                prediksi_statuscss: status_class,
                            },
                        ];
                    }
                }
            } else {
                loader_msg = json.message;
            }
            buttonLoading_class = "btn btn-primary"
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
        } else {
            if(msg_error != ""){
                isModalNotif = true
            }
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    function generate(field){
        let numbergenerate = (Math.floor(Math.random() * 10000) + 10000).toString().substring(1);
        switch(field){
            case "prediksi_field":
                prediksi_field = numbergenerate
                break;
        }
    }
    initapp();
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
{#if akses_page == true}
    <div class=" mx-auto px-2  ">
        <div class="flex gap-2">
            <div class="bg-white shadow-lg  w-[500px] ">
                <div class="flex flex-col w-full">
                    <div class="border-slate-200">
                        <div class="flex justify-start items-center border-b-2 p-2">
                            <h1 class="text-slate-600 font-bold text-sm lg:text-xl uppercase w-full">Prediksi</h1>
                            <div class="flex justify-end w-full p-2">
                                <button on:click={() => {
                                    callPrediksi();
                                }} class="btn btn-primary hover:bg-primary m-0 p-2 h-1 min-h-[40px]  rounded-md shadow-lg shadow-[#b3e4fc] border-none">Check</button>
                            </div>
                        </div>
                    </div>
                    <div class="w-full mt-2 p-2 flex flex-col gap-5 h-[500px]">
                        <select bind:value={prediksi_select_field} class="select select-bordered rounded-sm w-full ">
                            <option disabled selected value="">--Pilih Pasaran--</option>
                            {#each listHome as rec }
                                <option value="{rec.home_idcomp}">{rec.home_nama}</option>
                            {/each}
                        </select>
                        <div class="form-control w-full ">
                            <div class="input-group w-full">
                                <Input_custom
                                    input_required={true}
                                    input_tipe="number_nolabel"
                                    bind:value={prediksi_field}
                                    input_maxlenght="4"
                                    input_id="prediksi_field"
                                    input_placeholder="Prize 1"/>
                                <button on:click={() => {
                                    generate("prediksi_field");
                                }} class="btn btn-primary">Generator</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="bg-white shadow-lg p-2 w-full">
                <div class="flex flex-col w-full">
                    <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[550px] overflow-y-scroll">
                        <table class="table table-compact w-full ">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="1%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">STATUS</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">INVOICE</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">CODE</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">TANGGAL</th>
                                    <th width="*" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">USERNAME</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">TIPE</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-left">PERMAINAN</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-center">NOMOR</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">BET</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">DISC</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">KEI</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">BAYAR</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WIN</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-xs lg:text-sm text-white text-right">WIN TOTAL</th>
                                </tr>
                            </thead>
                            {#if listresult != ""}
                                <tbody>
                                    {#each listresult as rec}
                                    <tr>
                                        <td class="text-xs lg:text-sm align-top text-center">
                                            <span class="{rec.prediksi_statuscss} text-center rounded-md p-1 px-2 shadow-lg ">{rec.prediksi_status}</span>
                                        </td>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.prediksi_invoice}</td>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.prediksi_code}</td>
                                        <td class="text-xs lg:text-sm align-top text-center whitespace-nowrap">{rec.prediksi_tanggal}</td>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.prediksi_username}</td>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.prediksi_posisitogel}</td>
                                        <td class="text-xs lg:text-sm align-top text-left whitespace-nowrap">{rec.prediksi_permainan}</td>
                                        <td class="text-xs lg:text-sm align-top text-center font-semibold whitespace-nowrap">{rec.prediksi_nomor}</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.prediksi_bet)}</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.prediksi_diskon)}  ({rec.prediksi_diskonpercen}%)</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.prediksi_keipercen)}  ({rec.prediksi_keipercen}%)</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.prediksi_bayar)}</td>
                                        <td class="text-xs lg:text-sm align-top text-right font-semibold ">{new Intl.NumberFormat().format(rec.prediksi_win)}x</td>
                                        <td class="text-xs lg:text-sm align-top text-right text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.prediksi_totalwin)}</td>
                                    </tr>
                                    {/each}
                                </tbody>
                            {:else}
                                <tbody>
                                    <tr>
                                        <td colspan="16" class="text-center">
                                            <progress class="self-start progress progress-primary w-56"></progress>
                                        </td>
                                    </tr>
                                </tbody>
                            {/if}
                        </table>
                    </div>
                    <div class="bg-[#F7F7F7] rounded-sm h-20 p-2 ">
                        <table class="w-full">
                            <tr>
                                <td class="text-left font-semibold text-xs">TOTAL BET</td>
                                <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(totalbet)}</td>
                            </tr>
                            <tr>
                                <td class="text-left font-semibold text-xs">MEMBER WINLOSE</td>
                                <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(totalwin)}</td>
                            </tr>
                            <tr>
                                <td class="text-left font-semibold text-xs">COMPANY WINLOSE</td>
                                <td class="text-right text-xs {temp_companywinlose_class}">{new Intl.NumberFormat().format(subtotal)}</td>
                            </tr>
                        </table>
                    </div>
                </div>
            </div>
        </div>
        
    </div>
{/if}