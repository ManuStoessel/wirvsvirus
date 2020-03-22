<template>
    <div>

        <AboutUs></AboutUs>
        <b-row>
            <b-col>
                <Ticker></Ticker>
            </b-col>
            <b-col>
                <b-card
                        img-alt="Image"
                        img-top
                        tag="article"
                        class="mb-2 dealer"
                        style="max-width: 60%; text-align: left"
                >
                    <b-card-title title="Fabian Lauterbach"> </b-card-title>
                    <b-card-img :src="require('@/assets/LadenliebePerson.jpg')"></b-card-img>
                    <b-card-text>
                        <p style="text-align: left">"Ladenliebe ist das was wir jetzt brauchen. Gemeinsam können wir beweisen, dass wir eine
                            Gemeinschaft sind, die zusammenhält. Zeigen wir es dem Virus und machen LadenLiebe groß!"</p>
                    </b-card-text>
                </b-card>
            </b-col>
        </b-row>
        <hr>
        <h3>Wer sind wir?</h3>
        <b-row>
            <b-col cols="7">
                <b-embed class="videoEmbed"
                        type="iframe"
                        aspect="16by9"
                        src="https://www.youtube.com/embed/tqubXyCSWTQ"
                ></b-embed>
            </b-col>
            <b-col class="fastDonation" cols="4">
                <SearchDonateCity></SearchDonateCity>
            </b-col>
        </b-row>
        <hr>
        <LocalDealers :dealers="dealers"></LocalDealers>
    </div>
</template>

<script>
    import LocalDealers from "./LocalDealers";
    import SearchDonateCity from "./SearchDonateCity";
    import Ticker from "./Ticker";
    import AboutUs from "./AboutUs";



    export default {
        name: "Mainpage",
        components: {
            LocalDealers,
            SearchDonateCity,
            Ticker,
            AboutUs,
        },
        methods:{

        },
        props: {
            dealers: Array
        },
        created() {
        }
    }
</script>

<style scoped>
    h3{
        text-align: left;
    }
    .fastDonation{
        margin:auto;
    }
    .videoEmbed {
        margin: auto;
        min-width: 40%;
    }
</style>