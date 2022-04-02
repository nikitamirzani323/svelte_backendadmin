<script>
  import Router from "svelte-spa-router";
	import { wrap } from "svelte-spa-router/wrap";
  import Navigation from '../src/components/Navigation.svelte'
  import Sidebar from '../src/components/Sidebar.svelte'
  import NotFound from '../src/pages/NotFound.svelte'
  import Login from '../src/pages/Login.svelte'
  import Home from '../src/pages/Home.svelte'
  export let path_api = "";
  let token = localStorage.getItem("token");
	let routes = "";
	let isNav = false;
  let drawer_class = "drawer-mobile"
  drawer_class = ""
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
				},
				component: Home,
			}),
			"*": NotFound,
		};
  }
</script>

<main class="flex space-x-2 h-screen w-full bg-[#f0f2f5]">
  <div class="drawer {drawer_class}">
    <input id="my-drawer-2" type="checkbox" class="drawer-toggle">
    <div class="drawer-content flex flex-col gap-6">
      {#if isNav}
        <Navigation />
      {/if}
      <div class="container mx-auto px-2">
        {#if isNav}
        <div class="bg-white shadow-lg p-5 mt-5 mb-10">
          <Router {routes} />
        </div>
        {:else}
          <Router {routes} />
        {/if}
      </div>
      
    </div> 
    {#if isNav}
      <Sidebar />
    {/if}
  </div>
  
</main>

<style global lang="postcss">
  @tailwind base;
  @tailwind components;
  @tailwind utilities;
</style>
