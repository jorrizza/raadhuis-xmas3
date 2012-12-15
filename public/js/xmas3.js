var color = '#ff0000';

$(function () {
    $('#color-container button').click(
      function () {
        color = rgb2hex($(this).css('background-color'));
        $(".led").css({stroke: color});
        $.ajax(
          {
            url: '/set-color',
            data: {value: color.substring(1, 7)}
          }
        );
      }
    );
  }
);
 
function rgb2hex(rgb) {
  rgb = rgb.match(/^rgb\((\d+),\s*(\d+),\s*(\d+)\)$/);
  function hex(x) {
    return ("0" + parseInt(x).toString(16)).slice(-2);
  }
  return "#" +hex(rgb[1]) + hex(rgb[2]) + hex(rgb[3]);
}
