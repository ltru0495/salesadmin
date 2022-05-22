$(function() {
    // $('#infotable').DataTable({
    //     "language": {
    //         "lengthMenu": "Mostrar _MENU_ registro",
    //         "zeroRecords": "No se encontraron registros",
    //         "info": "Mostrando página _PAGE_ de _PAGES_",
    //         "infoEmpty": "No hay información disponible",
    //         "infoFiltered": "(Filtrado de _MAX_ en total)",
    //         "search": "Buscar",
    //         "paginate": {
    //             "first": "Primero",
    //             "last": "Último",
    //             "next": "Siguiente",
    //             "previous": "Anterior"
    //         },
    //     },
    //     "iDisplayLength": 50
    // });

    $('.changepass-btn').on('click', function() {
        console.log("ADASDS")
        $('#user').val(this.id);
        $('#pModal').modal('show');
    })


    $('#savepass').on('click', function(event) {
        event.preventDefault();
        username = $('#user').val();
        oldPass = $('#password').val();
        newPass = $('#newpassword').val();
        confPass = $('#confirmpassword').val();

        if (oldPass == "" || newPass == "") {
            alert("Debe llenar todos los campos");
            return;
        }
        if (confPass != newPass) {
        	alert("Las contraseñas no coinciden");
        	return;

        }

		var data = {
            username: username,
            oldPassword: oldPass,
            newPassword: newPass
        };
        $.ajax({
            type: "PUT",
            url: url_server + "/usuario/pass",
            data: data,
            success: function() {
                console.log("ASDADASDSSSSSSSSS")
                window.location.reload();
            },
        }).done(function() {
            console.log("ASDADASD")
            window.location.reload();
        });;


    })
});

$(document).ajaxStop(function() {
    window.location.reload();
});