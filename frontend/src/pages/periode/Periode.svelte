<script>
    import Home from "../periode/Home.svelte";
    export let path_api = ""
    let listHome = [];
    let listPeriodePasaran = [];
    let record = "";
    let totalrecord = 0;
    let token = localStorage.getItem("token");
    let akses_page = false;
    async function initapp() {
        const res = await fetch(path_api+"api/home", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "PERIODE-VIEW",
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
        listHome = [];
        listPeriodePasaran = [];
        const res = await fetch(path_api+"api/allperiode", {
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
            totalrecord = record.length;
            let recordpasaran = json.pasaranonline;
            let status_class = "";
            let totalmember_class = "";
            let totalbet_class = "";
            let totalbayar_class = "";
            let totalcancel_class = "";
            let winlose_class = "";
            let revisi_class = "";
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    if (record[i]["pasaran_status"] == "RUNNING") {
                        status_class = "bg-[#FFEB3B] text-black"
                    }else if(record[i]["pasaran_status"] == "DONE") {
                        status_class = "bg-[#8BC34A] text-black"
                    }else{
                        status_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_totalmember"]) > 0) {
                        totalmember_class = "text-blue-700";
                    } else {
                        totalmember_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_totalbet"]) > 0) {
                        totalbet_class = "text-blue-700";
                    } else {
                        totalbet_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_totaloutstanding"]) > 0) {
                        totalbayar_class = "text-blue-700";
                    } else {
                        totalbayar_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_totalcancelbet"]) > 0) {
                        totalcancel_class = "text-blue-700";
                    } else {
                        totalcancel_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_winlose"]) > 0) {
                        winlose_class = "text-blue-700";
                    } else {
                        winlose_class = "text-red-500";
                    }
                    if (parseInt(record[i]["pasaran_revisi"]) > 0) {
                        revisi_class = "text-blue-700";
                    } else {
                        revisi_class = "text-red-500";
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_no: record[i]["pasaran_no"],
                            home_invoice: record[i]["pasaran_invoice"].toString(),
                            home_idcompp: record[i]["pasaran_idcompp"],
                            home_code: record[i]["pasaran_code"],
                            home_periode: record[i]["pasaran_periode"],
                            home_name: record[i]["pasaran_name"],
                            home_keluaran: record[i]["pasaran_keluaran"],
                            home_tanggal: record[i]["pasaran_tanggal"],
                            home_totalmember:record[i]["pasaran_totalmember"],
                            home_totalmember_class: totalmember_class,
                            home_totalbet: record[i]["pasaran_totalbet"],
                            home_totalbet_class: totalbet_class,
                            home_totaloutstanding:record[i]["pasaran_totaloutstanding"],
                            home_totaloutstanding_class: totalbayar_class,
                            home_totalcancelbet:record[i]["pasaran_totalcancelbet"],
                            home_totalcancelbet_class:totalcancel_class,
                            home_msgrevisi: record[i]["pasaran_msgrevisi"],
                            home_revisi: record[i]["pasaran_revisi"],
                            home_revisi_class: revisi_class,
                            home_winlose: record[i]["pasaran_winlose"],
                            home_winlose_class: winlose_class,
                            home_status: record[i]["pasaran_status"],
                            home_status_class: status_class,
                        },
                    ];
                }
                if (recordpasaran != null) {
                for (var j = 0; j < recordpasaran.length; j++) {
                    listPeriodePasaran = [
                        ...listPeriodePasaran,
                        {
                            pasarancomp_idcompp: recordpasaran[j]["pasarancomp_idcompp"],
                            pasarancomp_nama: recordpasaran[j]["pasarancomp_nama"],
                        },
                    ];
                }
            }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleRefreshData = (e) => {
        listHome = [];
        totalrecord = 0;
        setTimeout(function () {
            initHome();
        }, 1000);
    };
    const handleLogout = (e) => {
        logout()
    }
    initapp();
</script>
{#if akses_page == true}
    <Home
        on:handleRefreshData={handleRefreshData}
        on:handleLogout={handleLogout}
        {path_api}
        {token}
        {listPeriodePasaran}
        {listHome}
        {totalrecord}/>
{/if}