<script>
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../components/Input.svelte' 
    export let path_api = "";
    let client_ipaddress = "";
    let client_timezone = "";
    const schema = yup.object().shape({
        username: yup.string().required().matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").max(30,"Username must be at most 30 characters"),
        password: yup.string().required().min(4,"Password must be at least 4 characters").max(50,"Password must be at most 50 characters")
    });
    const { form, errors, handleChange, handleSubmit } = createForm({
        initialValues: {
            username: "",
            password: "",
        },
        validationSchema: schema,
        onSubmit:(values) => {
            handleSave(values.username,values.password)
        }
    })
    async function handleSave(username,password) {
        const res = await fetch(path_api+"api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                username: username,
                password: password,
                ipaddress: client_ipaddress,
                timezone: client_timezone,
            }),
        });
        const json = await res.json();
        if (json.status == 400) {
            alert(json.message);
            username = "";
            password = "";
        } else {
            localStorage.setItem("token", json.token);
            localStorage.setItem("master", username);
            window.location.href = "/";
        }
    }
    async function initTimezone() {
        const res = await fetch(path_api+"api/healthz");
        if (!res.ok) {
            const message = `An error has occured: ${res.status}`;
            throw new Error(message);
        } else {
            const json = await res.json();
            client_ipaddress = json.real_ip;
            client_timezone = "Asia/Jakarta";
        }
    }
    initTimezone();
   
   
</script>
<section class="bg-white shadow-lg p-5 mt-5 mb-10 mx-[550px]">
    <div class="flex flex-col gap-4">
        <div class="space-y-4">
            <h1 class="text-center text-2xl font-semibold text-gray-500">LOGIN AGEN</h1>
        </div>
        <div class="relative form-control">
            <Input_custom
                input_onchange="{handleChange}"
                input_autofocus={true}
                input_required={true}
                input_tipe="text"
                input_invalid={$errors.username.length > 0}
                bind:value={$form.username}
                input_id="username"
                input_placeholder="Username"
                />
            {#if $errors.username}
                <small class="text-pink-600 text-[11px]">{$errors.username}</small>
            {/if}
        </div>
        <div class="relative form-control">
            <Input_custom
                input_onchange="{handleChange}"
                input_autofocus={false}
                input_required={true}
                input_tipe="password"
                input_attr="password"
                input_invalid={$errors.password.length > 0}
                bind:value={$form.password}
                input_id="password"
                input_placeholder="Password"
                />
            {#if $errors.password}
                <small class="text-pink-600 text-[11px]">{$errors.password}</small>
            {/if}
        </div>
        <div class="form-control">
            <button
                on:click={() => {
                    handleSubmit();
                }} 
                class="btn btn-primary rounded-md">Submit</button>
        </div>
    </div>
    
</section>
<style>
    .loaderbox{
        -webkit-box-reflect: below 1px linear-gradient(transparent, #0004);
    }
</style>
