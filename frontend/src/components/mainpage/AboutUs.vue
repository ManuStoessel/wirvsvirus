<template>
    <div>
        <div class="aboutUs">
            <b-row style="margin-right: 0">
                <b-col cols="7" class="aboutUsText">
                    <div style="font-size: 1.5rem; text-align: left">
                        <p>Werde zum Stadtretter und spende für deinen Lieblingsladen.</p>
                    </div>
                    <div style="font-size: 1.5rem; text-align: left">
                        <p>LadenLiebe bietet schnelle und direkte Hilfe für den Einzelhandel.</p>
                    </div>
                    <div style="font-size: 1.5rem; text-align: left">
                        <p>Sichere Abwicklung über PayPal.</p>
                    </div>
                </b-col>
            </b-row>
        </div>
        <div class="aboutUsAbout">
            <h2>An wen kannst du spenden?</h2>

            <p>Du kannst deinen Lieblingsladen direkt unterstützen (sei es ein Restaurant, eine Kneipe, eine Boutique, Blumenladen, Thaimassage, Kiosk, Friseur, etc.).</p>

            <p>Suche den Namen des Ladens oder grenze deine Suche über die Postleitzahl ein und lasse dir alle Läden deiner Region anzeigen.
                Sollte dein Lieblingsladen noch nicht gelistet sein kannst du einen Einladungslink an deren E-Mail-Adresse verschicken.</p>

            <h2>Du kannst die Einzelhändler deiner Stadt unterstützen.</h2>
            <p>Wähle hierzu einfach deine Stadt aus und entscheide, ob du eine bestimmte Branche unterstützen möchtest.
                Die Gewerbevereine , IHK und andere lokale Partner verantworten die Verteilung der Gelder und berücksichtigen
                hierbei, wer von der Krise am meisten betroffen ist.</p>
            <p>Nähere Informationen zu unseren Partnern findest du in der entsprechenden Rubrik.</p>
        </div>
        <hr>

    </div>
</template>

<script>
    export default {
        name: "AboutUs"
    }
</script>

<style scoped>
    .aboutUs{
        background-image: url('../../assets/LadenLiebeLogo.png');
        background-color: #cccccc;
        height: 500px;
        background-position: center;
        background-repeat: no-repeat;
        background-size: cover;
        width:100%;
        background-color: white;
        height: 100%;
        color:#f9f9f9;
    }
    .aboutUsImage{
        margin-right: 20%;
        margin-top: 3%;
        margin-bottom: 6%;

    }
    .aboutUsText{
        margin-top:2%;
        padding-left: 5%;
        padding-right: 10%;
        margin-bottom: 5%;
    }

    .aboutUsAbout{
        text-align: left;
    }

    .video{
        text-align: left;
    }
</style>