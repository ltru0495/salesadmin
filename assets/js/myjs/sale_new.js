$(function () {

    $('#infotable').DataTable({
        "language": {
            "lengthMenu": "",
            "zeroRecords": "No se encontraron registros",
            "info": "",
            "infoEmpty": "No hay información disponible",
            "infoFiltered": "(Filtrado de _MAX_ en total)",
            "search": "Buscar",
            "paginate": {
                "first": "",
                "last": "",
                "next": "",
                "previous": ""
            },
        },
        "iDisplayLength": 10000
    });
     $('#refundsTable').DataTable({
        "language": {
            "lengthMenu": "",
            "zeroRecords": "No se encontraron registros",
            "info": "",
            "infoEmpty": "No hay información disponible",
            "infoFiltered": "(Filtrado de _MAX_ en total)",
            "search": "Buscar",
            "paginate": {
                "first": "",
                "last": "",
                "next": "",
                "previous": ""
            },
        },
        "iDisplayLength": 10000
    });
    $('.paginate_button').css('display', 'none')
    $("form")[0].reset();

	$('#productButton').on( 'click' , (event) => {
		event.preventDefault();
		var code = $("#code").val();
		// console.log(code)
		code = code.toUpperCase();
		if(code === "") { 	
			alert("Codigo no especificado")
		} else {
			// console.log('/api/producto/'+code)
			$(':input','#saleform').not(':button, :submit, :reset, :hidden, #place').val('')
 
			$.getJSON(url_server+'/api/producto/'+code ,  function(product) { 
				$('#code').val(product.code);
				$('#brand').val(product.brand);
				$('#serie').val(product.serie);
				$('#size').val(product.size);
				$('#model').val(product.model);
				$('#location').val(product.location);
				$('#date').val(product.regdate);
				$('#sprice').val(product.sprice);
				$('#pnote').val(product.note);
				

				// console.log(product);
			}).fail(function() {
				alert( "PRODUCTO NO ENCONTRADO" );
			});

		}
	});
	// 000001MA22004
	// 000002MA21002
	// 000001MA22005

	let pm =""

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

	$('#seller').on('input', function (){
		$('#seller').val($('#seller').val().toUpperCase());
	});

	function getField(fieldName) {
		return($('#'+fieldName).val());
	}

	function setModalField(fieldName){
		$('#modal-'+fieldName).text(getField(fieldName));
		if(fieldName === "payment_method")  $('#modal-'+fieldName).text(pm);
		
		if(fieldName === "price") $('#modal-'+fieldName).text("S/ "+getField(fieldName));
	}


	function getObjectToSend() {
		var sale = {
			code: getField('code'),
			seller: getField('seller'),
			price: getField('price'),
			place: getField('place'),
			location: getField('location'),
			comment: getField('comment'),
			payment_method: pm,
			pnote: getField('pnote'),
		}
		return sale;
	}

	$('#save').on('click', (event) => {
	 	event.preventDefault();
	 	var sale = getObjectToSend();
	 	if(sale.code === "") {
	 		alert("Producto no especificado")
	 		return;
		}
		if (pm === ""){
			alert("Metodo de pago no seleccionado")
			return;
		}

	 	$.ajax({
		  type: "POST",
		  url: "/venta/nueva",
		  data: sale,
		  success: function(){
		  	// window.location.href = url_server+'/venta/nueva';
		  	window.location.reload();

		  },
		});
	});



	$("form").submit(function( event ) {
		
		event.preventDefault();

		setModalField('code');
		setModalField('brand');
		setModalField('serie');
		setModalField('size');
		setModalField('model');
		setModalField('location');
		setModalField('seller');
		setModalField('price');
		setModalField('payment_method');
		setModalField('place');
		setModalField('pnote');
		setModalField('comment');

		if (pm === ""){
			alert("Metodo de pago no seleccionado")
			return;
		}
		
		if(isNaN(Number($("#price").val()))) {
	 	    event.preventDefault();
 			alert("Precio debe ser un numero");
 		    return
		}
	    $('#confirmationModal').modal('show');
	    return;
	});
});
