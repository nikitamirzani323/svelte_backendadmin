<script>
    import { createForm } from "svelte-forms-lib";
    import * as yup from "yup";
    import Input_custom from '../components/Input.svelte' 
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
        }
    }
   
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
                input_value={$form.username}
                input_id="username"
                input_placeholder="Username"
                />
               
        </div>
        <div class="relative form-control">
            <Input_custom
                input_onchange="{handleChange}"
                input_autofocus={false}
                input_required={true}
                input_tipe="password"
                input_attr="password"
                input_invalid={$errors.password.length > 0}
                input_value={$form.password}
                input_id="password"
                input_placeholder="Password"
                />
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
