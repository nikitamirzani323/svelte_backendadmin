<script>
	import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
	import Navigation from '../src/components/Navigation.svelte'
	import NotFound from '../src/pages/NotFound.svelte'
	import Login from '../src/pages/Login.svelte'
	import Home from '../src/pages/home/Home.svelte'
	import Admin from '../src/pages/admin/Admin.svelte'
	import Adminrule from '../src/pages/adminrule/Adminrule.svelte'
	import Log from '../src/pages/log/Log.svelte'
	import Pasaran from '../src/pages/pasaran/Pasaran.svelte'
	import Periode from '../src/pages/periode/Periode.svelte'
	import Prediksi from '../src/pages/prediksi/Prediksi.svelte'
	import Report from '../src/pages/report/Report.svelte'
	import Tailwindcss from './Tailwindcss.svelte'
	export let path_api = "";
	export let font_size = "";
  	let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
  	if (token == "" || token == null) {
    // isNav = true;
		routes = {
			"/": wrap({
        		props: {
					path_api: path_api,
				},
				component: Login,
			}),
			"*": NotFound,
		};
	} else {
		isNav = true;
			routes = {
				"/": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Home,
				}),
			"/admin": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Admin,
				}),
			"/adminrule": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Adminrule,
				}),
			"/log": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Log,
				}),
			"/pasaran": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Pasaran,
				}),
			"/periode": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Periode,
				}),
			"/prediksi": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Prediksi,
				}),
			"/report": wrap({
					props: {
						path_api: path_api,
						font_size: font_size,
					},
					component: Report,
				}),
			"*": NotFound,
			};
  }
</script>
{#if isNav}
<main class="flex flex-col  h-screen w-full bg-gradient-to-bl from-[#EBF8FF] to-[#FFFFF0]">
	<Navigation />
	<div class="justify-center ">
    	<div class="w-full  mt-5">
      		<Router {routes} />
    	</div>
  	</div>
</main>
{:else}
<div class="container mx-auto">
  <Router {routes} />
</div>
{/if}
<Tailwindcss />