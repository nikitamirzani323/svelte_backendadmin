<script>
    import { createEventDispatcher } from "svelte";
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../../components/Input.svelte' 
    import Modal_popup from '../../components/Modal_popup.svelte'
    import Modal_alert from '../../components/Modal_alert.svelte' 
    import Loader from '../../components/Loader.svelte' 
    import Panel from '../../components/Panel_default.svelte' 
    import Panel_table from '../../components/panel_table.svelte' 

    export let path_api = "";
    export let token = "";
    export let listHome = [];
    export let listPeriodePasaran = [];
    export let totalrecord = 0;

    const schema = yup.object().shape({
        field_revisi: yup.string().required()
                    .matches(/^[a-zA-z0-9 ]+$/, "Revisi must Character A-Z or a-z or 1-9 or space ")
                    .min(5,"Revisi must be at least 5 characters")
                    .max(70,"Revisi must be at most 70 characters"),
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            field_revisi: "",
        },
        validationSchema: schema,
        onSubmit:(values) => {
            revisiTransaksi(values.field_revisi)
        }
    })

    let page = "Periode";
    let sData = "New";
    let isModal_Form_New = false
    let isModal_Form_Revisi = false
    let isModal_Form_MemberlistBet = false
    let isModal_Form_listBet = false
    let isModal_Form_listBetall = false
    let isModalLoading = false
    let isModalNotif = false
    let modal_width = "max-w-xl"
    let modal_width_form_revisi = "max-w-xl"
    let modal_width_listbetmember = "max-w-xl"
    let modal_width_listmembernomor = "max-w-xl"
    let modal_width_listbetall = "max-w-xl"
    let loader_class = "hidden"
    let loader_msg = "Sending..."
    let buttonLoading_class = "btn btn-primary"
    let msg_error = "";
    let searchMember = "";
    let searchMemberListBet = "";
    let searchListAllBet = "";
    let searchHome = "";
    let filterHome = [];
    let filterMember = [];
    let filterMemberListBet = [];
    let filterListBetALl = [];
    let listMember = [];
    let listBetUsername = [];
    let listBetTable = [];
    let listBetTableGroup = [];
    let listMemberNomor = [];
    let listBet = [];
    

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
    let temp_totalbayar = 0
    let temp_totalwinestimate = 0
    let temp_grandtotal = 0
    let temp_grandtotal_class = ""
    let client_username = ""
    let chooce_permainan = "";
    let select_pasaran = "";
    let tab_listmember = "bg-sky-600 text-white"
    let tab_betgroup = ""
    let tab_listbet = ""
    let panel_listmember = true
    let panel_betgroup = false

    let tab_listbetall_all = "bg-sky-600 text-white"
    let tab_listbetall_winner = ""
    let tab_listbetall_cancel = ""
    let panel_listbetall_all = true
    let panel_listbetall_winner = false
    let panel_listbetall_cancel = false

  

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
                EditData(idtrxkeluaran,idpasarancode,pasaran_msgrevisi)
            }
        } else {
            if(msg_error != ""){
                isModalNotif = true
            }
        }
    }
    async function SaveNewTransaksi() {
        let flag = false;
        msg_error = "";
        if (select_pasaran == "") {
            flag = true;
            msg_error += "The Pasaran is required<br>";
        }
        if (flag == false) {
            loader_class = "inline-block"
            loader_msg = "Sending..."
            isModalLoading = true;
            const res = await fetch(path_api+"api/savepasarannew", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    pasaran_code: select_pasaran,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message;
            } else {
                loader_msg = json.message;
            }
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            isModalLoading = false
            RefreshHalaman();
            select_pasaran = "";
        } else {
            if(msg_error != ""){
                isModalLoading = false;
                isModalNotif = true;
            }
        }
    }
    const NewData = () => {
        sData = "New";
        clearField()
        modal_width = "max-w-xl";
        isModal_Form_New = true;
    };
    async function EditData(e,y,z) {
        if(e != ""){
            isModalLoading = true;
            modal_width = "max-w-5xl";
            sData = "Edit";
            clearField();
            idtrxkeluaran = e;
            idpasarancode = y;
            pasaran_msgrevisi = z;
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
            dispatch("handleLogout", "call");
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
            } 
        }
    }
    async function call_listbettable() {
        listBetTable = [];
        const res = await fetch(path_api+"api/periodelistbettable", {
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
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetTable = [
                        ...listBetTable,
                        {
                            permainan: record[i]["permainan"],
                        },
                    ];
                }
            }
        }
    }
    async function call_bettable(e) {
        listBetTableGroup = [];
        chooce_permainan = e;
        const res = await fetch(path_api+"api/periodebettable", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listBetTableGroup = [
                        ...listBetTableGroup,
                        {
                            bet_keluaran: record[i]["bet_keluaran"],
                            bet_totalbet: record[i]["bet_totalbet"],
                            bet_totalmember: record[i]["bet_totalmember"],
                        },
                    ];
                }
            }
        }
    }
    async function call_listmembernomor(nomor) {
        listMemberNomor = [];
        temp_totalbayar = 0
        temp_totalwinestimate = 0
        temp_grandtotal = 0
        temp_grandtotal_class = "text-red-500 font-semibold"
        const res = await fetch(path_api+"api/periodelistmemberbynomor", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: chooce_permainan,
                nomor: nomor,
            }),
        });
        const json = await res.json();
        let record = json.record;
        let nomember = 0;
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            if (record != null) {
                modal_width_listmembernomor = "max-w-6xl"
                isModal_Form_listBet = true;
                for (var i = 0; i < record.length; i++) {
                    nomember = nomember + 1;
                    temp_totalbayar = temp_totalbayar + parseInt(record[i]["member_bayar"])
                    temp_totalwinestimate = temp_totalwinestimate + parseInt(record[i]["member_winhasil"])
                    listMemberNomor = [
                        ...listMemberNomor,
                        {
                            member_no: nomember,
                            member_name: record[i]["member"],
                            member_posisitogel: record[i]["member_posisitogel"],
                            member_permainan: record[i]["member_permainan"],
                            member_nomor: record[i]["member_nomor"],
                            member_bet: record[i]["member_bet"],
                            member_disc: record[i]["member_disc"],
                            member_discpercen: record[i]["member_discpercen"].toFixed(2),
                            member_kei: record[i]["member_kei"],
                            member_keipercen: record[i]["member_keipercen"].toFixed(2),
                            member_bayar: record[i]["member_bayar"],
                            member_win: record[i]["member_win"].toFixed(2),
                            member_winhasil: record[i]["member_winhasil"],
                        },
                    ];
                }
                temp_grandtotal = parseInt(temp_totalbayar) - parseInt(temp_totalwinestimate)
                if(temp_grandtotal > 0){
                    temp_grandtotal_class = "text-blue-700 font-semibold"
                }
            } 
        }
    }
    async function call_listbet(e) {
        listBet = [];
        totalbet = 0;
        totalbayar = 0;
        totalwin = 0;
        temp_grandtotal_class = "text-red-500 font-semibold"
        const res = await fetch(path_api+"api/periodelistbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                permainan: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            if (record != null) {
                let status_class = "";
                if(totalwin > 0){
                    temp_grandtotal_class = "text-blue-700 font-semibold"
                }else{
                    temp_grandtotal_class = "text-blue-700 font-semibold"
                }
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["bet_status"] != "CANCEL") {
                        if(record[i]["bet_status"] == "LOSE"){
                            status_class = "bg-red-400 text-white"
                        }else if(record[i]["bet_status"] == "CANCEL"){
                            status_class = "bg-red-600 text-white"
                        }else if(record[i]["bet_status"] == "WINNER"){
                            status_class = "bg-[#8BC34A] text-black"
                        }else{
                            status_class = "bg-[#FFEB3B] text-black"
                        }
                        listBet = [
                            ...listBet,
                            {
                                bet_id: record[i]["bet_id"].toString(),
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
                }
            } 
        }
    }
    async function call_listbetstatus(e) {
        listBet = [];
        totalbet = 0;
        totalbayar = 0;
        totalwin = 0;
        temp_grandtotal_class = "text-red-500 font-semibold"
        const res = await fetch(path_api+"api/periodelistbetstatus", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                idinvoice: parseInt(idtrxkeluaran),
                status: e,
            }),
        });
        const json = await res.json();
        let record = json.record;
        totalbet = json.totalbet;
        totalbayar = json.subtotal;
        totalwin = json.subtotalwin;
        if (json.status === 400) {
            dispatch("handleLogout", "call");
        } else {
            if (record != null) {
                let status_class = "";
                if(totalwin > 0){
                    temp_grandtotal_class = "text-blue-700 font-semibold"
                }else{
                    temp_grandtotal_class = "text-blue-700 font-semibold"
                }
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["bet_status"] == "LOSE"){
                        status_class = "bg-red-400 text-white"
                    }else if(record[i]["bet_status"] == "CANCEL"){
                        status_class = "bg-red-600 text-white"
                    }else if(record[i]["bet_status"] == "WINNER"){
                        status_class = "bg-[#8BC34A] text-black"
                    }else{
                        status_class = "bg-[#FFEB3B] text-black"
                    }
                    listBet = [
                        ...listBet,
                        {
                            bet_id: record[i]["bet_id"].toString(),
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
            } 
        }
    }
    async function cancelbetTransaksi(e,f) {
        msg_error = "";
        loader_class = "inline-block"
        loader_msg = "Sending..."
        isModalLoading = true;
        const res = await fetch(path_api+"api/cancelbet", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                sData: sData,
                page: "PERIODE-SAVE",
                permainan: f,
                idinvoice: parseInt(idtrxkeluaran),
                idinvoicedetail: parseInt(e),
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            loader_msg = json.message;
        } else if (json.status == 403) {
            loader_msg = json.message;
        } else {
            loader_msg = json.message;
        }
        setTimeout(function () {
            loader_class = "hidden";
        }, 1000);
        isModalLoading = false
        call_listbet("4D");
        call_listmember();
    }
    function callrevisiTransaksi() {
        modal_width_form_revisi = "max-w-xl";
        isModal_Form_Revisi = true;
    }
    async function revisiTransaksi(e) {
        let flag = true;
        msg_error = "";
        if(idtrxkeluaran == ""){
            flag = false
            msg_error = "Anda tidak bisa melakukan revisi"
        }
        if(idpasarancode == ""){
            flag = false
            msg_error = "Anda tidak bisa melakukan revisi"
        }
        if(flag){
            loader_class = "inline-block"
            loader_msg = "Sending..."
            isModalLoading = true;
            periode_status_field = "LOCK";
            const res = await fetch(path_api+"api/saveperioderevisi", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sData: sData,
                    page: "PERIODE-SAVE",
                    idinvoice: parseInt(idtrxkeluaran),
                    msgrevisi: e,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                loader_msg = json.message;
            } else if (json.status == 403) {
                loader_msg = json.message;
            } else {
                loader_msg = json.message;
            }
            setTimeout(function () {
                loader_class = "hidden";
            }, 1000);
            isModalLoading = false
            RefreshHalaman();
            EditData(idtrxkeluaran,idpasarancode,pasaran_msgrevisi)
            isModal_Form_Revisi = false;
        }else{
            if(msg_error != ""){
                isModalNotif = true
                isModalLoading = false
            }
        }
    }
    const handleSelectPermainangroup = (event) => {
        listBetTableGroup = [];
        if (event.target.value != "") {
            listBetTableGroup = [];
            call_bettable(event.target.value);
        }
    };
    const handleSelectPermainan = (event) => {
        if (event.target.value != "") {
            call_listbet(event.target.value);
        } else {
            listBet = [];
        }
    };
    const groupMember = (nomor) => {
        listMemberNomor = [];
        call_listmembernomor(nomor);
    };
    const ChangeTabMenu = (e) => {
        totalbet = 0;
        totalbayar = 0;
        totalwin = 0;
        switch(e){
            case "menu_listmember":
                tab_listmember = "bg-sky-600 text-white"
                tab_betgroup = ""
                tab_listbet = ""
                panel_listmember = true
                panel_betgroup = false
                break;
            case "menu_betgroup":
                call_listbettable();
                tab_listmember = ""
                tab_listbet = ""
                tab_betgroup = "bg-sky-600 text-white"
                panel_listmember = false
                panel_betgroup = true
                break;
            case "menu_listbet":
                modal_width_listbetall = "max-w-7xl"
                isModal_Form_listBetall = true;
                call_listbettable();
                listBet = [];
                tab_listmember = ""
                tab_listbet = "bg-sky-600 text-white"
                tab_betgroup = ""
                panel_listmember = false
                panel_betgroup = false
                break;
        }
    }
    const ChangeTabMenuListBet = (e) => {
        switch(e){
            case "menu_listbet_all":
                listBet = [];
                tab_listbetall_all = "bg-sky-600 text-white"
                tab_listbetall_winner = ""
                tab_listbetall_cancel = ""
                panel_listbetall_all = true
                panel_listbetall_winner = false
                panel_listbetall_cancel = false
                break;
            case "menu_listbet_winner":
                call_listbetstatus("winner")
                tab_listbetall_all = ""
                tab_listbetall_winner = "bg-sky-600 text-white"
                tab_listbetall_cancel = ""
                panel_listbetall_all = false
                panel_listbetall_winner = true
                panel_listbetall_cancel = false
                break;
            case "menu_listbet_cancel":
                call_listbetstatus("cancel")
                tab_listbetall_all = ""
                tab_listbetall_winner = ""
                tab_listbetall_cancel = "bg-sky-600 text-white"
                panel_listbetall_all = false
                panel_listbetall_winner = false
                panel_listbetall_cancel = true
                break;
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
 
    function clearField(){
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
        listMember = [];
        listBetUsername = [];
        listBetTable = [];
        listBetTableGroup = [];
        tab_listmember = "bg-sky-600 text-white"
        tab_betgroup = ""
        tab_listbet = ""
        panel_listmember = true
        panel_betgroup = false
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
        if (searchListAllBet) {
            filterListBetALl = listBet.filter(
                (item) =>
                    item.bet_status
                        .toLowerCase()
                        .includes(searchListAllBet.toLowerCase()) || 
                    item.bet_id
                        .toLowerCase()
                        .includes(searchListAllBet.toLowerCase()) ||
                    item.bet_username
                        .toLowerCase()
                        .includes(searchListAllBet.toLowerCase()) ||
                    item.bet_nomortogel
                        .toLowerCase()
                        .includes(searchListAllBet.toLowerCase())
            );
        } else {
            filterListBetALl = [...listBet];
        }
    }
</script>
<Loader loader_class="{loader_class}" loader_msg="{loader_msg}" />
<Panel
    on:handleNewForm={NewData}
    on:handleRefresh={RefreshHalaman}
    panel_button_new={true}
    panel_button_refresh={true}
    panel_page="{page}"
    panel_total="{totalrecord}">
    <slot:template slot="panel_search">
        <div class="absolute inset-y-0 left-0 flex items-center pl-2">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 stroke-current text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                <path stroke-linecap="round" stroke-linejoin="round" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
        </div>
        <input 
            bind:value={searchHome}
            type="text" placeholder="Search by Invoice, Status, Pasaran" class="input input-bordered w-full max-w-full rounded-md pl-8 pr-4 focus:ring-0 focus:outline-none">
    </slot:template>
    <slot:template slot="panel_body">
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
                            EditData(rec.home_invoice,rec.home_code,rec.home_msgrevisi);
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
    </slot:template>
</Panel>


<input type="checkbox" id="my-modal-formnew" class="modal-toggle" bind:checked={isModal_Form_New}>
<Modal_popup
    modal_popup_id="my-modal-formnew"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none w-11/12 {modal_width} overflow-hidden">
    <slot:template slot="modalpopup_body">
        {#if sData=="New"}
            <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2 ">
                <div class="relative form-control mt-2">
                    <select
                        class="select select-bordered w-full"
                        bind:value={select_pasaran}>
                        <option disabled selected value="">--Pilih Pasaran--</option>
                        {#each listPeriodePasaran as rec}
                            <option value={rec.pasarancomp_idcompp}>{rec.pasarancomp_nama}</option>
                        {/each}
                    </select>
                </div>
            </div>
            <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
                <button
                    on:click={() => {
                        SaveNewTransaksi();
                    }} 
                    class="{buttonLoading_class}">Submit</button>
            </div>
        {/if}
        {#if sData=="Edit"}
            <div class="flex justify-between  gap-2">
                <div class="w-1/2">
                    <div class="flex flex-auto flex-col overflow-auto gap-5 mt-2  ">
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
                    {:else if periode_statusrevisi_field == "OPEN"}
                        <div class="flex flex-wrap justify-end align-middle  mt-2">
                            <button
                                on:click={() => {
                                    callrevisiTransaksi();
                                }}  
                                class="btn btn-warning btn-block">Revisi</button>
                        </div>
                    {/if}
                    {#if pasaran_msgrevisi != ""}
                    <div class="alert alert-warning shadow-lg mt-2 rounded-sm">
                        <div>
                            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
                            <span class="text-sm">Alasan Revisi : {pasaran_msgrevisi}</span>
                        </div>
                    </div>
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
                        <li on:click={() => {
                                ChangeTabMenu("menu_listbet");
                            }}
                            class="items-center {tab_listbet} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">List Bet</li>
                    </ul>
                    {#if panel_listmember}
                        <h2 class="text-lg font-bold mb-2">List Member</h2>
                        <input 
                            bind:value={searchMember}
                            type="text" placeholder="Search by Username" class="input input-bordered w-full max-w-full rounded-md  focus:ring-0 focus:outline-none">
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
                    {/if}
                    {#if panel_betgroup}
                        <h2 class="text-lg font-bold mb-2">Bet Group</h2>
                        <select
                            on:change={handleSelectPermainangroup}
                            class="select select-bordered select-sm w-full rounded-sm focus:ring-0 focus:outline-none">
                            <option value="">--Pilih Permainan--</option>
                            {#each listBetTable as rec}
                                <option value={rec.permainan}>{rec.permainan}</option>
                            {/each}
                        </select>
                        <div class="w-full  scrollbar-thin scrollbar-thumb-sky-300 scrollbar-track-sky-100 h-[400px] overflow-y-scroll mt-2">
                            <table class="table table-compact w-full">
                                <thead class="sticky top-0">
                                    <tr>
                                        <th class="bg-[#6c7ae0] text-white text-xs text-center align-top">NOMOR</th>
                                        <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL MEMBER</th>
                                        <th class="bg-[#6c7ae0] text-white text-xs text-right align-top">TOTAL BET</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {#if listBetTableGroup  != ""}
                                        {#each listBetTableGroup  as rec}
                                            <tr>
                                                <td on:click={() => {
                                                    groupMember(rec.bet_keluaran);
                                                }} class="text-xs text-center align-top underline cursor-pointer">{rec.bet_keluaran}</td>
                                                <td class="text-xs text-right align-top text-blue-700 font-semibold">{rec.bet_totalmember}</td>
                                                <td class="text-xs text-right align-top text-blue-700 font-semibold">{new Intl.NumberFormat().format(rec.bet_totalbet)}</td>
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
                    {/if}
                </div>
            </div>
        {/if}
    </slot:template>
</Modal_popup>

<input type="checkbox" id="my-modal-formrevisi" class="modal-toggle" bind:checked={isModal_Form_Revisi}>
<Modal_popup
    modal_popup_id="my-modal-formrevisi"
    modal_popup_title="Entry/{sData}"
    modal_popup_class="select-none w-11/12 {modal_width_form_revisi} overflow-hidden">
    <slot:template slot="modalpopup_body">
        <div class="flex flex-auto flex-col overflow-auto gap-1 mt-2 ">
            <Input_custom
                input_onchange="{handleChange}"
                input_autofocus={true}
                input_required={true}
                input_tipe="text"
                input_invalid={$errors.field_revisi.length > 0}
                bind:value={$form.field_revisi}
                input_id="field_revisi"
                input_placeholder="Revisi"/>
            {#if $errors.field_revisi}
                <small class="text-pink-600 text-[11px]">{$errors.field_revisi}</small>
            {/if}
                
        </div>
        <div class="flex flex-wrap justify-end align-middle p-[0.75rem] mt-2">
            <button
                on:click={() => {
                    handleSubmit();
                }} 
                class="{buttonLoading_class}">Submit</button>
        </div>
    </slot:template>
</Modal_popup>

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
                type="text" placeholder="Search by Status,Code" class="input input-bordered mt-1 w-full max-w-full rounded-md  focus:ring-0 focus:outline-none">
           
            <Panel_table
                panel_class="mt-2" 
                panel_height="550px">
                <slot:template slot="paneltable_body">
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
                                        <td class="text-xs text-left align-top whitespace-nowrap font-bold">{rec.bet_nomortogel}</td>
                                        <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                                        <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_diskon)} ({rec.bet_diskonpercen}%)</td>
                                        <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_kei)} ({rec.bet_keipercen}%)</td>
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
                </slot:template>
            </Panel_table>
        </div>
    </div>
</div>


<input type="checkbox" id="my-modal-listmembernomor" class="modal-toggle" bind:checked={isModal_Form_listBet}>
<div class="modal" >
    <div class="modal-box relative  w-11/12 {modal_width_listmembernomor}  rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-listmembernomor" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">INFORMATION</h3>
            </div>
            <Panel_table
                panel_class="" 
                panel_height="550px">
                <slot:template slot="paneltable_body">
                    <table class="table table-compact w-full">
                        <thead class="sticky top-0">
                            <tr>
                                <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIPE</th>
                                <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">PERMAINAN</th>
                                <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">NOMOR</th>
                                <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BET</th>
                                <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">DISC</th>
                                <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">KEI</th>
                                <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BAYAR</th>
                                <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN</th>
                                <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">ESTIMATE WIN<br>TOTAL</th>
                            </tr>
                        </thead>
                        <tbody>
                            {#if listMemberNomor != ""}
                                {#each listMemberNomor as rec}
                                    <tr>
                                        <td class="text-xs text-left align-top whitespace-nowrap">{rec.member_name}</td>
                                        <td class="text-xs text-left align-top whitespace-nowrap">{rec.member_posisitogel}</td>
                                        <td class="text-xs text-left align-top whitespace-nowrap">{rec.member_permainan}</td>
                                        <td class="text-xs text-left align-top whitespace-nowrap font-bold">{rec.member_nomor}</td>
                                        <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_bet)}</td>
                                        <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_disc)}({rec.member_discpercen}%)</td>
                                        <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_kei)}({rec.member_keipercen}%)</td>
                                        <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_bayar)}</td>
                                        <td class="text-xs text-right align-top font-semibold whitespace-nowrap">{rec.member_win}x</td>
                                        <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.member_winhasil)}</td>
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
                </slot:template>
            </Panel_table>
            <div class="bg-[#F7F7F7] rounded-sm h-20 p-2">
                <table class="w-full">
                    <tr>
                        <td class="text-left font-semibold text-xs">TOTAL BAYAR</td>
                        <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(temp_totalbayar)}</td>
                    </tr>
                    <tr>
                        <td class="text-left font-semibold text-xs">TOTAL ESTIMATE WIN</td>
                        <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(temp_totalwinestimate)}</td>
                    </tr>
                    <tr>
                        <td class="text-left font-semibold text-xs">GRANDTOTAL</td>
                        <td class="text-right text-xs {temp_grandtotal_class}">{new Intl.NumberFormat().format(temp_grandtotal)}</td>
                    </tr>
                </table>
            </div>
        </div>
    </div>
</div>

<input type="checkbox" id="my-modal-listbetall" class="modal-toggle" bind:checked={isModal_Form_listBetall}>
<div class="modal" >
    <div class="modal-box relative  w-11/12 {modal_width_listbetall}  rounded-none lg:rounded-lg p-2  overflow-hidden">
        <div class="flex flex-col items-stretch">
            <div class="h-8">
                <label for="my-modal-listbetall" class="btn btn-xs lg:btn-sm btn-circle absolute right-2 top-2">✕</label>
                <h3 class="text-xs lg:text-sm font-bold mt-1">LIST BET</h3>
            </div>
            <ul class="flex justify-center items-center gap-2">
                <li on:click={() => {
                        ChangeTabMenuListBet("menu_listbet_all");
                    }}
                    class="items-center {tab_listbetall_all}  px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">All</li>
                <li on:click={() => {
                        ChangeTabMenuListBet("menu_listbet_winner");
                    }}
                    class="items-center {tab_listbetall_winner} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">WINNER</li>
                <li on:click={() => {
                        ChangeTabMenuListBet("menu_listbet_cancel");
                    }}
                    class="items-center {tab_listbetall_cancel} px-2 py-1.5 text-xs lg:text-sm cursor-pointer rounded-md outline outline-1 outline-offset-1 outline-sky-600">CANCEL</li>
            </ul>
            <Panel_table
                panel_class="mt-2" 
                panel_height="550px">
                <slot:template slot="paneltable_body">
                    {#if panel_listbetall_all}
                        <div class="flex justify-start items-stretch gap-2 mb-2 w-full">
                            <div class="p-0 w-full">
                                <select
                                    on:change={handleSelectPermainan}
                                    class="select select-bordered select-sm rounded-sm w-full focus:ring-0 focus:outline-none">
                                    <option value="">--Pilih Permainan--</option>
                                    {#each listBetTable as rec}
                                        <option value={rec.permainan}>{rec.permainan}</option>
                                    {/each}
                                </select>
                            </div>
                            <div class="p-0 w-full">
                                <input 
                                    bind:value={searchListAllBet}
                                    type="text" placeholder="Search by Status, Code, Nomor" class="input input-bordered input-sm  rounded-sm w-full focus:ring-0 focus:outline-none">
                            </div>
                        </div>
                        <table class="table table-compact w-full">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">&nbsp;</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">CODE</th>
                                    <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">IPADDRESS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">DEVICE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIMEZONE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIPE</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">PERMAINAN</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">NOMOR</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BET</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">DISC</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">KEI</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BAYAR</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN<br>TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#if filterListBetALl  != ""}
                                    {#each filterListBetALl as rec}
                                        <tr>
                                            <td class="text-xs text-center align-top whitespace-nowrap ">
                                                {#if periode_keluaran_field == ""}
                                                    {#if rec.bet_status == "RUNNING"}
                                                        <svg on:click={() => {
                                                            cancelbetTransaksi(
                                                                rec.bet_id,rec.bet_typegame
                                                            );
                                                        }} xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 cursor-pointer" fill="none" viewBox="0 0 24 24" stroke="currentColor" stroke-width="2">
                                                            <path stroke-linecap="round" stroke-linejoin="round" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                                                        </svg>
                                                    {/if}
                                                {/if}
                                            </td>
                                            <td class="text-xs text-center align-top whitespace-nowrap">
                                                <span class="{rec.bet_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.bet_status}</span>
                                            </td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_id}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_username}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_ipaddress}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_device}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_timezone}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_posisitogel}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_typegame}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap font-bold">{rec.bet_nomortogel}</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_diskon)}({rec.bet_diskonpercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_kei)}({rec.bet_keipercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                            <td class="text-xs text-right align-top font-semibold whitespace-nowrap">{rec.bet_win}x</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
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
                    {/if}
                    {#if panel_listbetall_winner}
                        <input 
                            bind:value={searchListAllBet}
                            type="text" placeholder="Search by Status, Code, Nomor" class="input input-bordered input-sm  rounded-sm w-full">
                        <table class="table table-compact w-full mt-2">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">CODE</th>
                                    <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">IPADDRESS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">DEVICE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIMEZONE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIPE</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">PERMAINAN</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">NOMOR</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BET</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">DISC</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">KEI</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BAYAR</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN<br>TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#if filterListBetALl  != ""}
                                    {#each filterListBetALl as rec}
                                        <tr>
                                            <td class="text-xs text-center align-top whitespace-nowrap">
                                                <span class="{rec.bet_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.bet_status}</span>
                                            </td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_id}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_username}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_ipaddress}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_device}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_timezone}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_posisitogel}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_typegame}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap font-bold">{rec.bet_nomortogel}</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_diskon)}({rec.bet_diskonpercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_kei)}({rec.bet_keipercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                            <td class="text-xs text-right align-top font-semibold whitespace-nowrap">{rec.bet_win}x</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
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
                    {/if}
                    {#if panel_listbetall_cancel}
                        <input 
                            bind:value={searchListAllBet}
                            type="text" placeholder="Search by Status, Code, Nomor" class="input input-bordered input-sm  rounded-sm w-full">
                        <table class="table table-compact w-full mt-2">
                            <thead class="sticky top-0">
                                <tr>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-center align-top">STATUS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">CODE</th>
                                    <th width="*" class="bg-[#6c7ae0] text-white text-xs text-left align-top">USERNAME</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">IPADDRESS</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">DEVICE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIMEZONE</th>
                                    <th width="1%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">TIPE</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">PERMAINAN</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-left align-top">NOMOR</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BET</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">DISC</th>
                                    <th width="10%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">KEI</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">BAYAR</th>
                                    <th width="7%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN</th>
                                    <th width="20%" class="bg-[#6c7ae0] text-white text-xs text-right align-top">WIN<br>TOTAL</th>
                                </tr>
                            </thead>
                            <tbody>
                                {#if filterListBetALl  != ""}
                                    {#each filterListBetALl as rec}
                                        <tr>
                                            <td class="text-xs text-center align-top whitespace-nowrap">
                                                <span class="{rec.bet_status_class} text-center rounded-md p-1 px-2 shadow-lg ">{rec.bet_status}</span>
                                            </td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_id}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_username}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_ipaddress}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_device}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_timezone}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_posisitogel}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap">{rec.bet_typegame}</td>
                                            <td class="text-xs text-left align-top whitespace-nowrap font-bold">{rec.bet_nomortogel}</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bet)}</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_diskon)}({rec.bet_diskonpercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_kei)}({rec.bet_keipercen}%)</td>
                                            <td class="text-xs text-right align-top text-blue-700 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_bayar)}</td>
                                            <td class="text-xs text-right align-top font-semibold whitespace-nowrap">{rec.bet_win}x</td>
                                            <td class="text-xs text-right align-top text-red-500 font-semibold whitespace-nowrap">{new Intl.NumberFormat().format(rec.bet_totalwin)}</td>
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
                    {/if}
                </slot:template>
            </Panel_table>
            
            <div class="bg-[#F7F7F7] rounded-sm h-20 p-2">
                <table class="w-full">
                    <tr>
                        <td class="text-left font-semibold text-xs">TOTAL BET</td>
                        <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(totalbet)}</td>
                    </tr>
                    <tr>
                        <td class="text-left font-semibold text-xs">TOTAL BAYAR</td>
                        <td class="text-right text-xs text-blue-700 font-semibold">{new Intl.NumberFormat().format(totalbayar)}</td>
                    </tr>
                    <tr>
                        <td class="text-left font-semibold text-xs">TOTAL WIN</td>
                        <td class="text-right text-xs {temp_grandtotal_class}">{new Intl.NumberFormat().format(totalwin)}</td>
                    </tr>
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
