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
				// console.log(product);
			}).fail(function() {
				alert( "PRODUCTO NO ENCONTRADO" );
			});

		}
	});

	$('#seller').on('input', function (){
		$('#seller').val($('#seller').val().toUpperCase());
	});

	function getField(fieldName) {
		return($('#'+fieldName).val());
	}
	// 000008MA21001

	function setModalField(fieldName){
		$('#modal-'+fieldName).text(getField(fieldName));
		if(fieldName === "payment_method")  $('#modal-'+fieldName).text($('#'+fieldName+" option:selected").text());
		
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
			payment_method: getField('payment_method'),
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
		setModalField('comment');
		
		if(isNaN(Number($("#price").val()))) {
	 	    event.preventDefault();
 			alert("Precio debe ser un numero");
 		    return
		}
	    $('#confirmationModal').modal('show');
 	    event.preventDefault();
	    return;
	});
});
