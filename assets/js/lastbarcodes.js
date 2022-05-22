$(function () {
	$.getJSON(url_server+'/api/barcodes',  function(products) { 
		products.forEach((product) => {
			for (var i = 0 ; i < 3 * product.quantity; i ++ ) {
				appendBarcodeSVG(product.code, i);
			}
		});
		
		
	}).fail(function() {
		alert( "PRODUCTOS NO REGISTRADOS HOY" );
	});


	function appendBarcodeSVG(code, index) {
		$('#barcodes').append(`<div class="col-md-4">
			<svg id="barcode_`+code+`_`+index+`"></svg>
			</div>`);
		JsBarcode("#barcode_"+code+"_"+index, code , {
			width : 64
		});
	}

})