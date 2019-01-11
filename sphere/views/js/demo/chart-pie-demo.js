// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';

// Pie Chart Example
var ctx = document.getElementById("myPieChart");
var myPieChart

$.get("/sim/by-type", function (data, status) {
  config = {
    type: "pie",
    data: {
      datasets: [

      ],
      labels: Object.keys(data),
    },

    options: {
      responsive: true
    }
  }
  var newDataSet = {
    type: "pie",
    backgroundColor: [],
    data: [],
    label: "Usage Report " + data.length,
  };
  var colorNames = Object.keys(window.chartColors);

  Object.keys(data).forEach(function (key, index) {
    newDataSet.data.push(data[key]);
    var colorName = colorNames[index % colorNames.length];
    var newColor = window.chartColors[colorName];
    newDataSet.backgroundColor.push(newColor)

  })
  config.data.datasets.push(newDataSet);
  myPieChart = new Chart(ctx, config);
})