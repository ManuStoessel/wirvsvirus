<template>
    <div class="aboutUs">
        <b-row>
            <b-col cols="7" class="aboutUsText">
                <div style="font-size: 2rem; text-align: left">
                    <h1>Lokalliebe!</h1>
                    <h2>Ein Ergebnis des WirVsVirus Hackathon!</h2>
                </div>
                <div style="font-size: 1.2rem; text-align: left">
                    Wir stellen eine Spendeplattform für lokale Unternehmen zur Verfügung. <br>
                    Unterstütze jetzt dein Lieblingsrestaurant, -bücherladen, -café
                </div>

            </b-col>
            <b-col class="aboutUsImage">
                <b-img src="https://wirvsvirushackathon.org/wp-content/uploads/2020/03/Hackathon_slogan_wei%C3%9F.png" fluid></b-img>
            </b-col>
        </b-row>

    </div>
</template>

<script>
    export default {
        name: "AboutUs"
    }
</script>

<style scoped>
    .aboutUs{
        margin-top: 1%;
        margin-bottom: 1%;
        font-size: 1.5rem;
        background-color: #b7b7b7;
        color:#f9f9f9;
    }
    .aboutUsImage{
        margin-right: 20%;
        margin-top: 3%;
        margin-bottom: 6%;

    }
    .aboutUsText{
        margin-top:3%;
        padding-left: 10%;
        padding-right: 10%;
    }
</style>