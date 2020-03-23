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
                        class=""
                >
                    <b-card-title title="Fabian Lauterbach"> </b-card-title>
                    <b-card-img :src="require('@/assets/LadenliebePerson.jpg')"></b-card-img>
                    <b-card-text>
                        <p style="text-align: left">"Ladenliebe ist das was wir jetzt brauchen. Gemeinsam können wir beweisen, dass wir eine
                            Gemeinschaft sind, die zusammenhält. Zeigen wir es dem Virus und machen LadenLiebe groß!"</p>
                    </b-card-text>
                </b-card>
            </b-col>
            <b-col class="">
                <SearchDonateCity></SearchDonateCity>
            </b-col>
        </b-row>
        <p style="margin-top: 4em"></p>
        <h2 style="text-align: left">Wer sind wir?</h2>
        <b-row class="">
            <b-col>
                <b-embed class="videoEmbed"
                        type="iframe"
                        aspect="16by9"
                        src="https://www.youtube.com/embed/jhXq4NM27Rc"
                ></b-embed>
            </b-col>
        </b-row>
        <p style="margin-top: 6em"></p>
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
    b-card {
        max-width: 100%;
        text-align: left;
        min-width: 60%
    }
    @media (min-width: 992px) {
        mobileScale {
            max-width: 60%;
            text-align: left;
            min-width: 60%
        }
    }

</style>