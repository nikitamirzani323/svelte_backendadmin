<script>
	import Modulepasaran from "../home/listpasaran.svelte";
    export let path_api = ""
    export let font_size = "";
	import dayjs from "dayjs";
    let listpasaran = [];
    let listperiode = [];
    let listHome = [];
    let listagen = [];
	let listHomepasaran = [];
    let listagenpasaran = [];
    let record = "";
    let totalrecord = 0;
    let admin_username = "";
    let token = localStorage.getItem("token");
    let akses_page = false;
	let winlose_year = dayjs().format("YYYY")
	let winlose_year_1 = parseInt(winlose_year)-1;
	let text_chart = "WINLOSE "+winlose_year;
	let text_chart_pasaran = "WINLOSE PASARAN "+winlose_year;

	let select_year = winlose_year;

	const handleSelect = (event) => {
		// alert(event.target.value)
        initHome(event.target.value);
		text_chart = "WINLOSE "+event.target.value;
    };
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
        } else if (json.status == 403) {
            alert(json.message);
            akses_page = false;
        } else {
            akses_page = true;
			initDashboardPasaran()
            initHome(winlose_year);
            initHomePasaran(winlose_year);
        }
    }
    async function initHome(e) {
		listHome = []
		listagen = []
        const res = await fetch(path_api+"api/dashboardwinlose", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
				"year":e
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listHome = [
                        ...listHome,
                        {
                            home_nmagen: record[i]["dashboardwinlose_nmagen"],
                            home_detail: record[i]["dashboardwinlose_detail"],
                        },
                    ];
                }
				for (let j = 0; j < listHome.length; j++) {
					let totaldata = listHome[j].home_detail.length;
					let temp = [];
					for (let x = 0; x < totaldata; x++) {
						temp.push(
							parseInt(listHome[j].home_detail[x]["dashboardwinlose_winlose"])/1000
						);
					}
					listagen = [
						...listagen,
						{
							name: listHome[j].home_nmagen,
							data: temp,
						},
					];
				}
            }
        } 
		createChart()
    }
	async function initDashboardPasaran() {
		listHomepasaran = []
		listagenpasaran = []
        const res = await fetch(path_api+"api/dashboardlistpasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({}),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
					let tutup = dayjs().format("DD MMM YYYY ")+record[i]["jamtutup"];
                    let jadwal = dayjs().format("DD MMM YYYY ")+record[i]["jamjadwal"];
                    let open = dayjs().format("DD MMM YYYY ")+record[i]["jamopen"];

                    listpasaran = [
                        ...listpasaran,
                        {
                            home_idcomppasaran: record[i]["idcomppasaran"],
                            home_codepasarantogel: record[i]["codepasarantogel"],
                            home_periode_idinvoice: record[i]["periode_idinvoice"],
                            home_periode_nomorperiode: record[i]["periode_nomorperiode"],
                            home_periode_tglperiode: record[i]["periode_tglperiode"],
                            home_periode_total_bayar: record[i]["periode_total_bayar"],
                            home_periode_total_bet: record[i]["periode_total_bet"],
                            home_periode_total_member: record[i]["periode_total_member"],
                            home_nmpasaran: record[i]["nmpasarantogel"],
                            home_pasaranurl: record[i]["pasaranurl"],
                            home_diundi: record[i]["pasarandiundi"],
                            home_jadwal: dayjs(jadwal).format("HH:mm"),
                            home_open: dayjs(open).format("HH:mm"),
                            home_tutup: dayjs(tutup).format("HH:mm"),
                        },
                    ];
                }
				
            }
        } 
    }
	async function initHomePasaran(e) {
		listHomepasaran = []
		listagenpasaran = []
        const res = await fetch(path_api+"api/dashboardwinlosepasaran", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
				"year":e
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    listHomepasaran = [
                        ...listHomepasaran,
                        {
                            home_nmagen: record[i]["dashboardagenpasaran_nmpasaran"],
                            home_detail: record[i]["dashboardagenpasaran_detail"],
                        },
                    ];
                }
				for (let j = 0; j < listHomepasaran.length; j++) {
					let totaldata = listHomepasaran[j].home_detail.length;
					let temp = [];
					for (let x = 0; x < totaldata; x++) {
						temp.push(
							parseInt(listHomepasaran[j].home_detail[x]["dashboardagenpasaran_winlose"])/1000
						);
					}
					listagenpasaran = [
						...listagenpasaran,
						{
							name: listHomepasaran[j].home_nmagen,
							data: temp,
						},
					];
				}
            }
        } 
		createChartPasaran()
    }
	async function createChart() {
		Highcharts.chart("container", {
			chart: {
				type: "column",
			},
			title: {
				text: "",
			},
			subtitle: {
				text: "",
			},
			xAxis: {
				categories: [
					"Jan",
					"Feb",
					"Mar",
					"Apr",
					"May",
					"Jun",
					"Jul",
					"Aug",
					"Sep",
					"Oct",
					"Nov",
					"Dec",
				],
				crosshair: true,
			},
			yAxis: {
				min: 0,
				title: {
					text: "Rupiah",
				},
			},
			tooltip: {
				headerFormat:
					'<span style="font-size:10px">{point.key}</span><table>',
				pointFormat:
					'<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f}</b></td></tr>',
				footerFormat: "</table>",
				shared: true,
				useHTML: true,
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0,
				},
			},
			series: listagen,
		});
	}
	async function createChartPasaran() {
		Highcharts.chart("containerpasaran", {
			chart: {
				type: "column",
			},
			title: {
				text: "",
			},
			subtitle: {
				text: "",
			},
			xAxis: {
				categories: [
					"Jan",
					"Feb",
					"Mar",
					"Apr",
					"May",
					"Jun",
					"Jul",
					"Aug",
					"Sep",
					"Oct",
					"Nov",
					"Dec",
				],
				crosshair: true,
			},
			yAxis: {
				min: 0,
				title: {
					text: "Rupiah",
				},
			},
			tooltip: {
				headerFormat:
					'<span style="font-size:10px">{point.key}</span><table>',
				pointFormat:
					'<tr><td style="color:{series.color};padding:0">{series.name}: </td>' +
					'<td style="padding:0"><b>{point.y:.1f}</b></td></tr>',
				footerFormat: "</table>",
				shared: true,
				useHTML: true,
			},
			plotOptions: {
				column: {
					pointPadding: 0.2,
					borderWidth: 0,
				},
			},
			series: listagenpasaran,
		});
	}
    async function logout() {
		localStorage.clear();
		window.location.href = "/";
	}
	const handleRefreshDataListPasaran = (e) => {
		listpasaran = []
        setTimeout(function () {
            initDashboardPasaran();
        }, 1000);
    };
    initapp();
</script>
{#if akses_page == true}
<div class="bg-[#f0f2f5] pb-10">
	<div class="">
		
		<Modulepasaran
			on:handleRefreshDataListPasaran={handleRefreshDataListPasaran}
			{path_api}
			{token}
			{font_size}
			{listpasaran}/>
		

		<div class="bg-white shadow-lg p-5 ">
			<div class="flex flex-col gap-2">
				<div class="flex">
					<h1 class="text-slate-600 font-bold text-sm lg:text-2xl uppercase w-full">{text_chart}</h1>
					<div class="hidden sm:flex md:flex justify-end w-28">
						<select
							on:change={handleSelect}
							bind:value="{select_year}" 
							class="select select-bordered w-full rounded-sm focus:ring-0 focus:outline-none">
							<option disabled selected value="">--Pilih Year--</option>
							<option value="{winlose_year}">{winlose_year}</option>
							<option value="{winlose_year_1}">{winlose_year_1}</option>
						</select>
					</div>
				</div>
				<div class="hidden  sm:inline w-full scrollbar-thin scrollbar-thumb-sky-200 h-[500px] ">
					<div class="w-full h-[500px]" id="container" />
				</div>
			</div>
		</div>

		<div class="bg-white shadow-lg p-5 mt-2">
			<div class="flex flex-col gap-2">
				<div class="flex">
					<h1 class="text-slate-600 font-bold text-sm lg:text-2xl uppercase w-full">{text_chart_pasaran}</h1>
					<div class="hidden sm:flex md:flex justify-end w-28">
						
					</div>
				</div>
				<div class="hidden  sm:inline w-full scrollbar-thin scrollbar-thumb-sky-200 h-[500px] ">
					<div class="w-full h-[500px]" id="containerpasaran" />
				</div>
			</div>
		</div>
	</div>
</div>
{/if}