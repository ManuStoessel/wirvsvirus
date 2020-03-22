<template>
    <div class="timer">
        <b-card
                title="Bisherige Spenden"
                style="max-width: 20rem;"
                class="mb-2 timer"
        >
            <b-card-text style="font-size: 5rem">
                {{numberOfDonation}}
            </b-card-text>


            <b-list-group >
                <div style="text-align: left">Neusten Spenden</div>
                <b-list-group-item v-for="(donation, index) in this.newestDonations" :key="index" style="font-size: 0.8rem; text-align:left">
                    {{donation.user}}:
                    <br>
                    {{donation.text}}
                </b-list-group-item>
            </b-list-group>
        </b-card>
    </div>
</template>

<script>
    export default {
        name: "Ticker",
        data(){
            return{
                numberOfDonation: 0,
                newestDonations: [
                    {
                        user: "Olaf",
                        text: "Lasst uns gemeinsam unsere lokalen Geschäfte unterstützen"
                    },
                    {
                        user: "Maria",
                        text: "Ich kann zwar nicht einkaufen, dennoch kann ich meine Lieblingsläden unterstützen"
                    }
                ]
            }
        },
        created() {
            setInterval(this.updateNumberOfDonations, 1000);
            setInterval(this.updateNewestDonations, 5000);
        },
        methods:{
            updateNumberOfDonations: function(){
                this.numberOfDonation++;
            },
            updateNewestDonations: function(){

            }
        }
    }
</script>

<style scoped>
    .timer{
        margin:auto;
        align-content: center;
    }
</style>