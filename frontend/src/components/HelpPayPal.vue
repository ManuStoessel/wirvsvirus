<template>
    <div>
        <b-card title="PayPal Donate-Button" sub-title="Anleitung zur Erstellung">
            <b-card-text>
                <b-list-group>
                    <b-list-group-item>
                        Klicken Sie auf "Hier gehts zu PayPal". <br>
                        Dort klicken Sie auf "Button jetzt erstellen.<br>
                        Anschließend können Sie ihren Donate-Button konfigurieren.<br>
                        Am Ende können Sie Ihren Donate-Button kopieren.<br>
                        Wir Benötigten jedoch nur den, wie im Beispiel, markierten Wert.
                    </b-list-group-item>
                </b-list-group>
                    <pre>
                    Beispiel:
                    &lt;form action="https://www.paypal.com/cgi-bin/webscr" method="post" target="_top"&gt;
                        &lt;input type="hidden" name="cmd" value="_s-xclick" /&gt;
                        &lt;input type="hidden" name="hosted_button_id" value="<b><u><mark>AB2382JGPY5SS</mark></u></b>" /&gt;
                        &lt;input type="image" src="https://www.paypalobjects.com/de_DE/DE/i/btn/btn_donateCC_LG.gif" border="0"
                        name="submit" title="PayPal - The safer, easier way to pay online!" alt="Spenden mit dem PayPal-Button" /&gt;
                        &lt;img alt="" border="0" src="https://www.paypal.com/de_DE/i/scr/pixel.gif" width="1" height="1" /&gt;
                    &lt;/form&gt;
                    </pre>

            </b-card-text>
            <a href="https://www.paypal.com/de/cgi-bin/webscr?cmd=_donate-intro-outside" class="PayPayDonateLink">Hier gehts zu PayPal</a>
        </b-card>
    </div>

</template>

<script>
    export default {
        name: "HelpPayPal"
    }
</script>

<style scoped>

</style>