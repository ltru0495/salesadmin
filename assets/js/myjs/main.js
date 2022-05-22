// inventory_url = url_server +"/inventario";
// inventory_report_url = url_server + "/productos/reporte";
newproduct_url = url_server +"/producto/nuevo";
// sales_url = url_server+"/";
// sales_date_url = url_server + "/ventas";
new_sale_url = url_server +"/venta/nueva";
refund_url = url_server +"/ventas/devolver"

// if(window.location.href === inventory_url ||  window.location.href===newproduct_url  
// 	|| window.location.href===inventory_report_url ) {
// 	$('.inventory-collapse ').addClass('show');
// }

// if(window.location.href === sales_url ||  window.location.href=== sales_date_url
//  || window.location.href === new_sale_url || window.location.href===refund_url  ) {
// 	$('.sales-collapse').addClass('show');
// }

if(window.location.href === new_sale_url || window.location.href === newproduct_url || window.location.href === refund_url) {
	$('.form-control').css('background-color', "#eee");
	$('.form-control').css('color', 'black');
	$('.form-control').css('border', '2px solid black');
	$('.form-control[readonly]').css('background-color', '#1f95E0')
	$('.form-control[readonly]').css('color', '#fff');
}

$("#sidebarCollapse").on("click", function() {
	$(".sidebar").toggleClass("hidden");
	$(".main-panel").toggleClass("allwidth");
})