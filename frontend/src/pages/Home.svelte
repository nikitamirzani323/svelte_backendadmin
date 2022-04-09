<script>
    export let path_api = "";
    let token = localStorage.getItem("token");
    async function initapp() {
		const res = await fetch(path_api+"api/home", {
			method: "POST",
			headers: {
				"Content-Type": "application/json",
				Authorization: "Bearer " + token,
			},
			body: JSON.stringify({
				page: "DASHBOARD-VIEW",
			}),
		});
		const json = await res.json();
		if (json.status === 400) {
			logout();
		} 
	}
    async function logout() {
		localStorage.clear();
		window.location.href = "/";
	}
    initapp();
</script>