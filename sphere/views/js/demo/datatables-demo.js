// Call the dataTables jQuery plugin

dt = $('#dataTable').DataTable().data();

$.get("/sockets", function (data, status) {

  data.forEach(element => {
    dt.row.add([element.ID, element.Typex, element.Status]).draw(true)

  });

})
