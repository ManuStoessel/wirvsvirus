<template>
    <div>
        <AboutUs></AboutUs>
        <b-row>
            <b-col>
                <Ticker></Ticker>
            </b-col>
            <b-col class="videoEmbed">
                <b-embed
                        type="iframe"
                        aspect="16by9"
                        src="https://www.youtube.com/embed/xw4mcqIxmdY"
                ></b-embed>
            </b-col>
            <b-col class="fastDonation">
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
    .fastDonation{
        margin: auto;
    }
    .videoEmbed {
        margin: auto;
        min-width: 40%;
    }
</style>