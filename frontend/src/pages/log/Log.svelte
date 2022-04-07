<script>
    import Home from "../log/Home.svelte";
    export let path_api = ""
    let listHome = [];
    let record = "";
    let totalrecord = 0;
    let admin_username = "";
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
                page: "ADMIN-VIEW",
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
        const res = await fetch(path_api+"api/log", {
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
            let no = 0;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1
                    let note = record[i]["log_note"];
                    let note_result = note.replace("\n", "<br>");
                    listHome = [
                        ...listHome,
                        {
                            home_no: no,
                            home_id: record[i]["log_id"],
                            home_datetime: record[i]["log_datetime"],
                            home_username: record[i]["log_username"],
                            home_page: record[i]["log_page"],
                            home_tipe: record[i]["log_tipe"],
                            home_note: note_result,
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
        {listHome}
        {totalrecord}/>
{/if}