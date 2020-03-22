<template>
    <div>
        <div class="mx-auto" style="max-width: 75%">

            <b-img
            :src="company.images"
            alt="Unternehmensbild"
            fluid>
            </b-img>

            <div align="left" >
                <h1>{{company.companyName}}</h1>
                {{ company.description }}

                <div style="margin-top: 10px">
                    <h1>Kontakt</h1>
                    <p>Adresse: {{company.street}}, {{company.town}}</p>
                    <p>E-Mail: <a :href="'mailto:' + company.email">{{ company.email }} </a></p>
                </div>
            </div>

            <div align="left">
                <form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_blank">
                    <input type="hidden" name="cmd" value="_s-xclick" />
                    <input type="hidden" name="hosted_button_id" :value="company.payPalButtonId" />
                    <input type="image" src="https://www.paypalobjects.com/de_DE/DE/i/btn/btn_donateCC_LG.gif" border="0" name="submit" title="PayPal - Der sichere, einfache Weg online zu bezahlen!" alt="Mit dem PayPal Button spenden" />
                    <img alt="" border="0" src="https://www.paypal.com/de_DE/i/scr/pixel.gif" width="1" height="1" />
                </form>
            </div>
        </div>

    </div>
</template>

<script>
    export default {
        name: "Registration_Company",

        data() {
            return {

                text: 'default'
            }
        },
        methods: {
        },
        props: {
            company: Object
        }
    }
</script>

<style scoped>
    b-img {
        width: 90%;
        height: auto;
    }
</style>