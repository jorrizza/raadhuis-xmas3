function color_poller () {
  $.ajax(
    {
      url: '/get-color'
    }
  ).done(
    function (color) {
      $(".led").css({stroke: '#'+color});
    }
  );

  setTimeout('color_poller()', 500);
}

$(function () {
    color_poller();

    $('#color-container button').click(
      function () {
        var color = rgb2hex($(this).css('background-color'));

        $.ajax(
          {
            url: '/set-color',
            data: {value: color.substring(1, 7)}
          }
        );
      }
    );

    $('#color-container button').hover(
      function () {
        $(this).find('img').show();
      },
      function () {
        $(this).find('img').hide();
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
