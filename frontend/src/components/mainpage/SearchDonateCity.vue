<template>
    <div class="fastDonation">
        <b-card
                title="Unterstütze deine Stadt!!!"
                tag="article"
                style="max-width: 20rem;"
                class="mb-2"
        >
            <b-card-text>
                <b-row>
                    <b-col class="cityInputField">
                        <b-input v-model="city" placeholder="Suche deine Stadt">
                            <b-icon-x-circle-fill align="right"></b-icon-x-circle-fill>
                        </b-input>

                        <b-button variant="danger" class="sendButton">Jetzt helfen</b-button>
                    </b-col>
                </b-row>
                <b-row>
                    <b-col class="cityDonationText">
                        <h2>
                            Einfach spenden!
                        </h2>
                        <a style="text-align: left">
                            Das Geld wird dann von lokalen Gewerbevereinen an die Unternehmen verteilt, die es benötigen.
                        </a>
                    </b-col>
                </b-row>
            </b-card-text>
        </b-card>
    </div>
</template>

<script>
    import {BIconXCircleFill } from 'bootstrap-vue'

    export default {

        name: "SearchDonateCity",
        data(){
            return{
                city: ""
            }
        },
        methods:{
            getPossibleCities: function(){
                console.log("RestCall to Backend, to receive Possible Cities")
            }
        },
        watch:{
            city: function(){
                this.getPossibleCities();
            }
        },
        components: {
            BIconXCircleFill
        }
    }
</script>

<style scoped>
    .fastDonation{
        margin-top: 2%;
    }

    .cityInputField{
        width: 100%;
        margin: auto;
    }

    b-nput{
        width:100%;
    }

    h2{
        margin-bottom: 2%;
    }

    .cityDonationText{
        margin-right: 10%;
        text-align: left;
    }

    .sendButton{
        margin-top: 4%;
        float: left;
    }


</style>