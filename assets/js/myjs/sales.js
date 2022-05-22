$(function(){
	// $('#infotable').DataTable({
	// 	"language": {
	//         "lengthMenu": "Mostrar _MENU_ registro",
	//         "zeroRecords": "No se encontraron registros",
	//         "info": "Mostrando página _PAGE_ de _PAGES_",
	//         "infoEmpty": "No hay información disponible",
	//         "infoFiltered": "(Filtrado de _MAX_ en total)",
	//         "search" : "Buscar",
	//         "paginate": {
	// 	        "first":      "Primero",
	// 	        "last":       "Último",
	// 	        "next":       "Siguiente",
	// 	        "previous":   "Anterior"
	// 	    },
	//     },
	//     "iDisplayLength": 50
	// });

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

    $('.paginate_button').css('display', 'none')
	// $('#infotable_filter').css('float', 'right');
	// $('#infotable_info').css('display', 'inline-block');
	// $('#infotable_paginate').css('display', 'inline-block');
	// $('#infotable_paginate').css('float', 'right');


	

});