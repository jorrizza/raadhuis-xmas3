var color = '#ff0000';

$(function () {
    $.farbtastic('#color').linkTo(function (c) {
                                    color = c;
                                    $(".led").css({stroke: color});
                                  }
                                 ).setColor(color);
    $('#set-color').click(
      function () {
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
