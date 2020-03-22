<template>
    <div>
        <form>
            <label for="defaultFormLoginEmailEx" class="grey-text">Deine E-Mail</label>
            <input type="email" id="defaultFormLoginEmailEx" class="form-control"/>
            <br/>
            <label for="defaultFormLoginPasswordEx" class="grey-text">Dein Passwort</label>
            <input type="password" id="defaultFormLoginPasswordEx" class="form-control"/>
            <div class="text-center mt-4">
                <b-btn class="btn" type="submit" @click="send">Anmdelden</b-btn>
            </div>
        </form>
    </div>
</template>

<script>

    export default {
        name: "Login",
        methods: {
            send: function(){
                this.$emit("login")
            },
            hashPassword: function(){

            }
        },
        data(){
            return {
                username: "",
                password: ""
            }
        }
    }
</script>

<style scoped>
    .loginfield{
        margin-top: 1%;
        margin-left:auto;
        margin-right:auto;
        width: 70%;
    }
</style>