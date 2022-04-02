<script>
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    export let path_api = "";
    let client_ipaddress = "";
    let client_timezone = "";
    const schema = yup.object().shape({
        username: yup.string().required().matches(/^[a-zA-z0-9]+$/, "Username must Character A-Z or a-z or 1-9 ").max(30),
        password: yup.string().required().min(4).max(50)
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
    $:{
        if ($errors.username || $errors.password){
            alert($errors.username+"\n"+$errors.password)
            $form.username = ""
            $form.password = ""
        }
    }
</script>
<section class="bg-white shadow-lg p-5 mt-5 mb-10 mx-[550px]">
    <div class="flex flex-col gap-4">
        <div class="space-y-4">
            <h1 class="text-center text-2xl font-semibold text-gray-500">LOGIN AGEN</h1>
        </div>
        <div class="relative form-control">
            <input
                autofocus
                on:change="{handleChange}"
                bind:value={$form.username}
                invalid={$errors.username.length > 0}
                type="text" 
                id="username"
                name="username"
                placeholder="Username"
                autocomplete="off"
                class="peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent"> 
            <label for="username" class="absolute left-3 top-[-0.7rem] text-gray-600 text-sm cursor-text 
                    transition-all
                    peer-placeholder-shown:text-base 
                    peer-placeholder-shown:text-gray-400 
                    peer-placeholder-shown:top-3 
                    peer-focus:text-[#1a73e8]
                    peer-focus:bg-[#fff]
                    peer-focus:text-[.75rem]
                    ">Username*</label>
        </div>
        <div class="relative form-control">
            <input
                on:change="{handleChange}"
                bind:value={$form.password}
                invalid={$errors.password.length > 0}
                type="password" 
                id="password"
                name="password"
                placeholder="Password"
                autocomplete="off"
                class="peer w-full rounded px-3  border border-gray-300  focus:border-blue-700 focus:ring-1 focus:ring-blue-700 focus:outline-none input active:outline-none placeholder-transparent"> 
            <label for="password" class="absolute left-3 top-[-0.7rem] text-gray-600 text-sm cursor-text 
                    transition-all
                    peer-placeholder-shown:text-base 
                    peer-placeholder-shown:text-gray-400 
                    peer-placeholder-shown:top-3 
                    peer-focus:text-[#1a73e8]
                    peer-focus:bg-[#fff]
                    peer-focus:text-[.75rem]
                    ">Password*</label>
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
    *{
        font-family: 'Open Sans', sans-serif;
    }
    .input {
        transition: border 0.2s ease-in-out;
        min-width: 280px;
    }
    .input:focus+.label,
    .input:active+.label,
    .input.filled+.label{
        font-size: .75rem;
        transition: all 0.2s ease-out;
        top:-0.9rem;
        background-color: #fff;
        color:#1a73e8;
        padding:0 5px 0 5px;
        margin: 0 5px 0 5px;
    }
    .label {
        transition: all 0.2s ease-out;
        top:0.4rem;
        left: 0;
    }
</style>