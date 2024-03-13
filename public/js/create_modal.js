let modal = document.getElementById("create_modal")
let button = document.getElementById("create_modal_button")

function openCreateModal() {
    modal.style.display = "flex";
    setTimeout(() => {
        modal.style.opacity = 1
    })
}

function closeCreateModal() {
    modal.style.opacity = 0
    setTimeout(() => {
        modal.style.display = "none";
    }, 300)
    console.log("Close modal")
}

button.onclick = function () {
    if(modal.style.display === "flex") {
        closeCreateModal()
        return
    }
    openCreateModal()
}

window.onclick = function (event) {
    if(event.target.id === modal.id) {
        closeCreateModal()
    }
    if (!event.target.matches('.dropbtn')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}

function openDropdown() {
    document.getElementById("create_modal_dropdown").classList.toggle("show");
}
