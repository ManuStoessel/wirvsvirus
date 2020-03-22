<template>
    <div class="navbar">
        <b-navbar type="light" variant="light">
            <div class="navbar-collapse collapse w-100 order-1 order-md-0 dual-collapse2">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item active">
                        <a class="nav-link" href="#" @click="changeView('mainpage')">Startseite</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#" @click="changeView('search')">Suche</a>
                    </li>
                    <li class="nav-item" v-if="this.loggedIn">
                        <a class="nav-link" href="#">Dashboard</a>
                    </li>
                </ul>
            </div>
            <div class="mx-auto order-0">
                <button class="navbar-toggler" type="button" data-toggle="collapse" data-target=".dual-collapse2">
                    <span class="navbar-toggler-icon"></span>
                </button>
            </div>
            <div class="navbar-collapse collapse w-100 order-3 dual-collapse2" v-if="loggedIn">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        <a class="nav-link" href="#" @click="changeView('profile')">Profil</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#" @click="logOff">Abmelden</a>
                    </li>
                </ul>
            </div>
            <div class="navbar-collapse collapse w-100 order-3 dual-collapse2" v-if="!loggedIn">
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item">
                        <a class="nav-link" href="#" v-b-modal.modal-1>Anmelden</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" href="#" v-b-modal.modal-2>Registrieren</a>
                    </li>
                </ul>
            </div>
        </b-navbar>


        <b-modal hide-footer id="modal-1" title="Anmelden">
            <Login @login="loggedInFunc"></Login>
        </b-modal>

        <b-modal hide-footer id="modal-2" size="xl" title="Registrieren">
            <Registration_Company></Registration_Company>
        </b-modal>
    </div>
</template>

<script>
    import Login from "../login/Login";
    import Registration_Company from "../Registration_Company";

    export default {
        name: "Navbar",
        components: {
            Login,
            Registration_Company
        },
        data(){
            return{
                loggedIn: false
            }
        },
        props:{
            comp: String
        },
        methods:{
            loggedInFunc(){
                this.loggedIn = true;
            },
            logOff(){
                this.loggedIn = false;
            },
            changeView(v){
                this.$emit('changeView', v);
            }
        }
    }
</script>

<style scoped>
    .navbar{
        width: 100%;
        margin:auto;
    }
</style>