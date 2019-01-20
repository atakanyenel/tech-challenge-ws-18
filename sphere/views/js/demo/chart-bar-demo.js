// Set new default font family and font color to mimic Bootstrap's default styling
Chart.defaults.global.defaultFontFamily = '-apple-system,system-ui,BlinkMacSystemFont,"Segoe UI",Roboto,"Helvetica Neue",Arial,sans-serif';
Chart.defaults.global.defaultFontColor = '#292b2c';

// Bar Chart Example

$.get("/bar-chart",function(data,status){
var ctx = document.getElementById("myBarChart");
console.log("bar-chart")

newBarChart={
  type: 'bar',
  data: {
    labels: [],
    datasets: [{
      label: "Daily Usage",
      backgroundColor: "rgba(2,117,216,1)",
      borderColor: "rgba(2,117,216,1)",
      data: [],
    }],
  },
  options: {
    scales: {
      xAxes: [{
        time: {
          unit: 'month'
        },
        gridLines: {
          display: false
        },
        ticks: {
          maxTicksLimit: 6
        }
      }],
      yAxes: [{
        ticks: {
          min: 0,
          max: 100,
          maxTicksLimit: 5
        },
        gridLines: {
          display: true
        }
      }],
    },
    legend: {
      display: false
    }
  }
}
Object.keys(data["values"]).forEach(function(key,index){
  newBarChart.data.labels.push(key)
  newBarChart.data.datasets[0].data.push(data["values"][key])
})

var myLineChart = new Chart(ctx,newBarChart )

})

;
