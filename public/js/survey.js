function submitAll() {
    const forms = document.getElementsByClassName("question");
    for (let i = 0; i < forms.length; i++) {
        htmx.trigger('#' + forms[i].id, 'ballons');
    }
}
