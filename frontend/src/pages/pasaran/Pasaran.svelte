<script>
    import Home from "../pasaran/Home.svelte";
    export let path_api = ""
    let listHome = [];
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
                page: "PASARAN-VIEW",
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
            initPasaran();
        }
    }
    async function initPasaran() {
        const res = await fetch(path_api+"api/allpasaran", {
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
            if (record != null) {
                let status_class = "";
                let active_class = "";
                for (var i = 0; i < record.length; i++) {
                    if(record[i]["statuspasaran"] == "ONLINE"){
                        status_class = "bg-[#FFEB3B] text-black"
                    }else{
                        status_class = "bg-[#E91E63] text-white"
                    }
                    if(record[i]["statuspasaranactive"] == "ACTIVE"){
                        active_class = "bg-[#8BC34A] text-black"
                    }else{
                        active_class = "bg-[#E91E63] text-white"
                    }
                    listHome = [
                        ...listHome,
                        {
                            home_id: record[i]["idcomppasaran"],
                            home_nama: record[i]["nmpasarantogel"],
                            home_tipe: record[i]["tipepasaran"],
                            home_diundi: record[i]["pasarandiundi"],
                            home_jamtutup: record[i]["jamtutup"],
                            home_jamjadwal: record[i]["jamjadwal"],
                            home_jamopen: record[i]["jamopen"],
                            home_display: record[i]["displaypasaran"],
                            home_status: record[i]["statuspasaran"],
                            home_status_class: status_class,
                            home_active: record[i]["statuspasaranactive"],
                            home_active_class: active_class,
                        },
                    ];
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
            initPasaran();
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
        {listHome}
        {totalrecord}/>
{/if}