// This is probably all useless

function submitAll() {
    /** @type {[HTMLFormElement]} */
    const forms = document.getElementsByClassName("question");
    const totalData = new FormData();
    const ids = [];
    const answers = [];

    for (let i = 0; i < forms.length; i++) {
        let data = new FormData(forms[i]);
        console.log("Form " + i);

        if (!data.get("answer")) {
            console.error("Question " + (i + 1) + " not answered");
            return;
        }

        ids.push(data.get("questionId"));
        answers.push(data.get("answer"));
    }

    console.log(answers);
}

htmx.defineExtension("aiden-test", {
    // onEvent: function(name, evt) {
    //     console.log(`Name: |${name}|`);
    //     console.log(`Event: |${evt}|`);
    // },
    encodeParameters: function(xhr, parameters, elt) {
    },
});
