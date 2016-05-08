window.onload = function () {


  moment.locale('ru');
  // Set new thresholds
  moment.relativeTimeThreshold('s', 60);
  moment.relativeTimeThreshold('m', 60);
  moment.relativeTimeThreshold('h', 30);
  moment.relativeTimeThreshold('d', 100);
  moment.relativeTimeThreshold('M', 10);


  var WEDDING_DATE = moment([2016, 6, 2, 14]);
  var days_left = 60,
    hours_left = 24,
    minutes_left = 60,
    seconds_left = 60;

  var days_bar,
    hours_bar,
    minutes_bar,
    seconds_bar;

  days_bar = new ProgressBar.Circle(".date-days", {
    color: '#aaa',
    // This has to be the same size as the maximum width to
    // prevent clipping
    strokeWidth: 4,
    trailWidth: 1,
    easing: 'easeInOut',
    duration: 1400,
    text: {
      autoStyleContainer: false
    },
    from: { color: '#fff', width: 0.5 },
    to: { color: '#fff', width: 2 },
    // Set default step function for all animate calls
    step: function(state, circle) {
      circle.path.setAttribute('stroke', state.color);
      circle.path.setAttribute('stroke-width', state.width);

      var diff = moment.duration(days_left, "days").humanize().split(" ");
      var diff_str = diff[diff.length - 1];

      circle.setText('<span class="time-left">' + days_left + '</span>' + '<span class="time-str">' + diff_str + '</span>');
    }
  });
  days_bar.text.style.fontFamily = '"Raleway", Helvetica, sans-serif';
  days_bar.text.style.fontSize = '2rem';



  var hours_bar = new ProgressBar.Circle(".date-hours", {
    color: '#aaa',
    // This has to be the same size as the maximum width to
    // prevent clipping
    strokeWidth: 4,
    trailWidth: 1,
    easing: 'easeInOut',
    duration: 1400,
    text: {
      autoStyleContainer: false
    },
    from: { color: '#fff', width: 0.5 },
    to: { color: '#fff', width: 2 },
    // Set default step function for all animate calls
    step: function(state, circle) {
      circle.path.setAttribute('stroke', state.color);
      circle.path.setAttribute('stroke-width', state.width);

      var diff = moment.duration(hours_left, "hours").humanize().split(" ");
      var diff_str = diff[diff.length - 1];

      circle.setText('<span class="time-left">' + hours_left + '</span>' + '<span class="time-str">' + diff_str + '</span>');
    }
  });
  hours_bar.text.style.fontFamily = '"Raleway", Helvetica, sans-serif';
  hours_bar.text.style.fontSize = '2rem';



  var minutes_bar = new ProgressBar.Circle(".date-minutes", {
    color: '#aaa',
    // This has to be the same size as the maximum width to
    // prevent clipping
    strokeWidth: 4,
    trailWidth: 1,
    easing: 'easeInOut',
    duration: 1400,
    text: {
      autoStyleContainer: false
    },
    from: { color: '#fff', width: 0.5 },
    to: { color: '#fff', width: 2 },
    // Set default step function for all animate calls
    step: function(state, circle) {
      circle.path.setAttribute('stroke', state.color);
      circle.path.setAttribute('stroke-width', state.width);

      var value = minutes_left;

      var diff = moment.duration(value, "minutes").humanize().split(" ");
      var diff_str = diff[diff.length - 1];

      circle.setText('<span class="time-left">' + value + '</span>' + '<span class="time-str">' + diff_str + '</span>');
    }
  });
  minutes_bar.text.style.fontFamily = '"Raleway", Helvetica, sans-serif';
  minutes_bar.text.style.fontSize = '2rem';



  var seconds_bar = new ProgressBar.Circle(".date-seconds", {
    color: '#aaa',
    // This has to be the same size as the maximum width to
    // prevent clipping
    strokeWidth: 4,
    trailWidth: 1,
    easing: 'easeInOut',
    duration: 1400,
    text: {
      autoStyleContainer: false
    },
    from: { color: '#fff', width: 0.5 },
    to: { color: '#fff', width: 2 },
    // Set default step function for all animate calls
    step: function(state, circle) {
      circle.path.setAttribute('stroke', state.color);
      circle.path.setAttribute('stroke-width', state.width);

      var value = seconds_left;

      var diff = moment.duration(value, "seconds").humanize().split(" ");
      var diff_str = diff[diff.length - 1];

      circle.setText('<span class="time-left">' + value + '</span>' + '<span class="time-str">' + diff_str + '</span>');
    }
  });
  seconds_bar.text.style.fontFamily = '"Raleway", Helvetica, sans-serif';
  seconds_bar.text.style.fontSize = '2rem';


  function timer() {
    now = moment();
    days_left = WEDDING_DATE.diff(now, 'days');
    hours_left = WEDDING_DATE.diff(now, 'hours') % 24;
    minutes_left = WEDDING_DATE.diff(now, 'minutes') % 60;
    seconds_left = WEDDING_DATE.diff(now, 'seconds') % 60;

    if (hours_left === 0) {
      days_left -= 1;
      hours_left = 23;
    }

    if (days_left < 0) {
      days_left = 0;
    }

    days_bar.set(days_left / 60);
    hours_bar.set(hours_left / 24);
    minutes_bar.set(minutes_left / 60);
    seconds_bar.set(seconds_left / 60);
  }


  var i = setInterval(timer, 1000);

  //bar.animate(1.0);  // Number from 0.0 to 1.0



/*
  var bar = new ProgressBar.Path('#date-days-path', {
    easing: 'easeInOut',
    duration: 1400,
    text: {
      autoStyleContainer: false
    },
    step: function (state, circle) {
      circle.path.setAttribute('stroke', state.color);
      circle.path.setAttribute('stroke-width', state.width);

      var value = Math.round(circle.value() * 100);
      if (value === 0) {
        circle.setText('');
      } else {
        circle.setText(value);
      }

    }
  });

  bar.set(0);
  bar.animate(1.0);  // Number from 0.0 to 1.0
*/

  var source = '{{#guests}}<div class="col-md-3 bestfriends cs-style-3 wow fadeInUp animated" data-wow-delay="0.5s" style="visibility: visible; animation-delay: 0.5s; animation-name: fadeInUp;"> \
      <a target="_blank" href="https://vk.com/id{{id}}"> <img src="{{photo}}" alt="friends" class="img-responsive"></a> \
      <h4>{{name}}</h4> \
      <h5>{{title}}</h5> \
      </a> \
    </div> \
  {{/guests}}';

  VK.init({ apiId: 5453424, onlyWidgets: false });


  function notAdmins(value) {
    return value >= 10;
  }


  var all_guests = GUESTS.map(function (person) { return person.id; });

  var getByID = function (id) {
    for (var i = 0; i < GUESTS.length; i++) {
      if (id === GUESTS[i].id) {
        return GUESTS[i];
      }
    }
    return 0;
  };


  VK.Api.call('users.get', {user_ids: all_guests, fields: "photo_100,photo_200,photo_400,status"}, function(r) {
    if (!r.response) {
      return false;
    }

    //console.log("response:", r.response);

      var users_with_photo = r.response.map(function (person) {
        var user = getByID(person.uid)
        user.photo = person.photo_200 || person.photo_400 || person.photo_max || person.photo_100;

        return user;
      });

      var rendered = Mustache.render(source, { guests: users_with_photo });

      var guests_block = document.querySelector(".guests");
      guests_block.innerHTML = rendered;
  });


  VK.Widgets.Comments("vk_comments", {limit: 15, width: "400px", attach: "*"});

  var feed = new Instafeed({
        /*clientId: '0775682bcdf64195b06e3c3eefa916dd',*/
        accessToken: '3218923292.467ede5.12c6de4a3a884d23a49fa4c3ab8ab5ff',
        target: 'instafeed',
        get: 'user',
        userId: '3218923292',
        tagName: 'loveday27',
        links: true,
        limit: 8,
        sortBy: 'most-recent',
        resolution: 'low_resolution',
        template: '<div class="col-xs-12 col-sm-6 col-md-4 col-lg-3"><div class="photo-box"><div class="image-wrap"><a target="_blank" href="{{link}}"><img src="{{image}}"></a></div><div class="description">{{caption}}<div class="date">{{model.date}}</div></div></div></div>'
    });
  feed.run();

};


