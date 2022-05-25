$(function() {
    var seller, comment, price, pm="";

    $(".selected").css("background-color", "#3ebf4c")
	$("#efectivo").on("click", () => {
		pm = "efectivo"
		document.getElementById("efectivo").style.background ="#3ebf4c"
		document.getElementById("electronico").style.background ="#ecf0f5"
		
	})
	$("#electronico").on("click", () => {
		pm = "electronico"
		document.getElementById("electronico").style.background ="#3ebf4c"
		document.getElementById("efectivo").style.background ="#ecf0f5"
	})

    function setVariables() {
        seller = $('#seller').val();
        comment = $('#comment').val();
        price= $('#price').val();
    }

    $('#seller').on('input', function (){
        $('#seller').val($('#seller').val().toUpperCase());
    });


    $("form").submit(function(event) {
        event.preventDefault();
        setVariables();

        if (pm === ""){
			alert("Metodo de pago no seleccionado")
			return;
		}
        if ( seller == "" || price == "" || isNaN(price)){
            alert("Ha ocurrido un error");
            return
        }
        $('#confirmModal').modal("show");
    });

    $("#confirm").on('click', function(event) {
        event.preventDefault();
        setVariables();
        var password = $('#password').val();
        var realid = $('#realid').text();

        if (password === "") {
            alert("Debe ingresar la contraseña");
            return;
        }
        if (pm === ""){
			alert("Metodo de pago no seleccionado")
			return;
		}
        var data = {
            password: password,
            price: Number(price),
            seller: seller,
            payment_method: pm,
            comment: comment,
        };

        $.ajax({
            type: "PUT",
            url: url_server + "/venta/" + realid,
            data: data,
            success: function() {
                window.location.href = "/venta/" + realid;
            },
            fail: function() {
                window.location.href = "/venta/" + realid;
            },
            complete: function(response) {
                window.location.href = "/venta/" + realid;
            }
        }).done(function() {
            window.location.href = "/inventario"
        });;
    });

});