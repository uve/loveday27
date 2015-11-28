
 var App = {
      countries:  [
            ["AF","Afghanistan",1856781],
            ["AL","Albania",1798686],
            ["DZ","Algeria",6669927],
            ["AD","Andorra",71575],
            ["AO","Angola",4286821],
            ["AG","Antigua and Barbuda",81545],
            ["AR","Argentina",24973660],
            ["AM","Armenia",1300013],
            ["AW","Aruba",81945],
            ["AU","Australia",21176595],
            ["AT","Austria",7135168],
            ["AZ","Azerbaijan",5737223],
            ["BS","Bahamas",293875],
            ["BH","Bahrain",1297500],
            ["BD","Bangladesh",10867567],
            ["BB","Barbados",224588],
            ["BY","Belarus",4856969],
            ["BE","Belgium",9441116],
            ["BZ","Belize",90939],
            ["BJ","Benin",460232],
            ["BM","Bermuda",63987],
            ["BT","Bhutan",211896],
            ["BO","Bolivia",3970587],
            ["BA","Bosnia Herzegovina",2582502],
            ["BW","Botswana",268038],
            ["BR","Brazil",107822831],
            ["BN","Brunei",277589],
            ["BG","Bulgaria",4083950],
            ["BF","Burkina Faso",741888],
            ["BI","Burundi",146219],
            ["KH","Cambodia",828317],
            ["CM","Cameroon",1486815],
            ["CA","Canada",33000381],
            ["KY","Cayman Islands",47003],
            ["CF","Central African Republic",161524],
            ["TD","Chad",317197],
            ["CL","Chile",11686746],
            ["CN","China",641601070],
            ["CO","Colombia",25660725],
            ["KM","Comoros",49320],
            ["CG","Congo",87559],
            ["CK","Cook Islands",1378],
            ["CR","Costa Rica",2511139],
            ["HR","Croatia",2780534],
            ["CU","Cuba",3090796],
            ["CY","Cyprus",726663],
            ["CZ","Czech Republic",8322168],
            ["CI","Côte d'Ivoire",565874],
            ["DK","Denmark",5419113],
            ["DJ","Djibouti",80378],
            ["DM","Dominica",42735],
            ["DO","Dominican Republic",5072674],
            ["EC","Ecuador",6012003],
            ["EG","Egypt",40311562],
            ["SV","El Salvador",1742832],
            ["GQ","Equatorial Guinea",124035],
            ["ER","Eritrea",59784],
            ["EE","Estonia",1047772],
            ["ET","Ethiopia",1636099],
            ["FO","Faeroe Islands",43605],
            ["FJ","Fiji",325717],
            ["FI","Finland",5117660],
            ["FR","France",55429382],
            ["PF","French Polynesia",161025],
            ["GA","Gabon",168592],
            ["GM","Gambia",271711],
            ["GE","Georgia",2188311],
            ["DE","Germany",71727551],
            ["GH","Ghana",5171993],
            ["GI","Greece",6438325],
            ["GR","Greenland",39717],
            ["GD","Grenada",47903],
            ["GU","Guam",112196],
            ["GT","Guatemala",2716781],
            ["GN","Guinea",205194],
            ["GW","Guinea-Bissau",57764],
            ["GY","Guyana",295200],
            ["HT","Haiti",1217505],
            ["HN","Honduras",1602558],
            ["HK","Hong Kong SAR",5751357],
            ["HU","Hungary",7388776],
            ["IS","Iceland",321475],
            ["IN","India",243198922],
            ["ID","Indonesia",42258824],
            ["IR","Iran",22200708],
            ["IQ","Iraq",2707928],
            ["IE","Ireland",3817491],
            ["IL","Israel",5928772],
            ["IT","Italy",36593969],
            ["JM","Jamaica",1393381],
            ["JP","Japan",109252912],
            ["JO","Jordan",3375307],
            ["KZ","Kazakhstan",9850123],
            ["KE","Kenya",16713319],
            ["KI","Kiribati",12156],
            ["KR","South Korea",45314248],
            ["KW","Kuwait",3022010],
            ["KG","Kyrgyzstan",1359416],
            ["LV","Latvia",1560452],
            ["LB","Lebanon",3336517],
            ["LS","Lesotho",110065],
            ["LR","Liberia",190731],
            ["LY","Libya",1362604],
            ["LI","Liechtenstein",34356],
            ["LT","Lithuania",2113393],
            ["LU","Luxembourg",510177],
            ["MG","Madagascar",17321756],
            ["MW","Malawi",12150362],
            ["MY","Malaysia",675074],
            ["MV","Maldives",16645],
            ["ML","Mali",11862559],
            ["MT","Malta",173003],
            ["MH","Marshall Islands",1246],
            ["MQ","Martinique",303302],
            ["MR","Mauritania",455553],
            ["MU","Mauritius",76681],
            ["YT","Mayotte",107940],
            ["MX","Mexico",50923060],
            ["FM","Micronesia",29370],
            ["MD","Moldova",1550925],
            ["MC","Monaco",34214],
            ["MN","Mongolia",514254],
            ["ME","Montenegro",364978],
            ["MA","Morocco",20207154],
            ["MZ","Mozambique",1467687],
            ["MM","Myanmar",624991],
            ["NA","Namibia",347414],
            ["NP","Nepal",3411948],
            ["NL","Netherlands",16143879],
            ["NC","New Caledonia",163997],
            ["NZ","New Zealand",4162209],
            ["NI","Nicaragua",891675],
            ["NE","Niger",298310],
            ["NG","Nigeria",67101452],
            ["NU","Niue",617],
            ["NO","Norway",4895885],
            ["OM","Oman",2584316],
            ["PK","Pakistan",20073929],
            ["PA","Panama",1899892],
            ["PG","Papua New Guinea",187284],
            ["PY","Paraguay",2005278],
            ["PE","Peru",12583953],
            ["PH","Philippines",39470845],
            ["PL","Poland",25666238],
            ["PT","Portugal",7015519],
            ["PR","Puerto Rico",2027549],
            ["QA","Qatar",2191866],
            ["RO","Romania",11178477],
            ["RU","Russia",84437793],
            ["RW","Rwanda",1110043],
            ["WS","Samoa",26977],
            ["SM","San Marino",16631],
            ["ST","Sao Tome and Principe",48806],
            ["SA","Saudi Arabia",17397179],
            ["SN","Senegal",3194190],
            ["RS","Serbia",4705141],
            ["SC","Seychelles",50220],
            ["SL","Sierra Leone",92232],
            ["SG","Singapore",4453859],
            ["SK","Slovakia",4507849],
            ["SI","Slovenia",1501039],
            ["SB","Solomon Islands",43623],
            ["SO","Somalia",163185],
            ["ZA","South Africa",24909854],
            ["ES","Spain",35010273],
            ["LK","Sri Lanka",4267507],
            ["SD","Sudan",9307189],
            ["SR","Suriname",201963],
            ["SZ","Swaziland",301211],
            ["SE","Sweden",8581261],
            ["CH","Switzerland",7180749],
            ["SY","Syria",5860788],
            ["TJ","Tajikistan",1357400],
            ["TZ","Tanzania",7590794],
            ["TH","Thailand",19386154],
            ["TL","Timor-Leste",11472],
            ["TG","Togo",319822],
            ["TO","Tonga",40131],
            ["TT","Trinidad and Tobago",856544],
            ["TN","Tunisia",5053704],
            ["TR","Turkey",35358888],
            ["TM","Turkmenistan",424855],
            ["TV","Tuvalu",3768],
            ["UG","Uganda",6523949],
            ["UA","Ukraine",16849008],
            ["AE","United Arab Emirates",8807226],
            ["GB","United Kingdom",57075826],
            ["US","United States",279834232],
            ["UY","Uruguay",2017280],
            ["UZ","Uzbekistan",11914665],
            ["VU","Vanuatu",29791],
            ["VE","Venezuela",14548421],
            ["VN","Viet Nam",39772424],
            ["YE","Yemen",4778488],
            ["ZM","Zambia",2313013],
            ["ZW","Zimbabwe",2852757]],

      langs: {"Mandarin Chinese": ["China", "Singapore"],
              "Hindi": ["India", "Fiji"],
              "Spanish": ["Argentina", "Bolivia", "Chile", "Colombia", "Costa Rica", "Cuba", "Dominican Republic", "Ecuador", "El Salvador", "Equatorial Guinea", "Guatemala", "Honduras", "Mexico", "Nicaragua", "Panama", "Paraguay", "Peru", "Spain", "United States", "New Mexico", "Puerto Rico", "Uruguay", "Venezuela"],
              "English": ["Antigua and Barbuda", "Australia", "The Bahamas", "Bangladesh", "Barbados", "Belize", "Botswana", "Brunei", "Cameroon", "Canada", "Dominica", "Ethiopia", "Eritrea", "Fiji", "The Gambia", "Ghana", "Grenada", "Guyana", "Hong Kong", "India", "Ireland", "Jamaica", "Kenya", "Kiribati", "Lesotho", "Liberia", "Malawi", "Maldives", "Malta", "Marshall Islands", "Maritius", "Micronesia", "Namibia", "Nauru", "New Zealand", "Nigeria", "Pakistan", "Palau", "Papua New Guinea", "Philippines", "Rwanda", "Saint Kitts and Nevs", "Saint Lucia", "Saint Vincent and the Grenadines", "Samoa", "Seychelles", "Sierra Leone", "Singapore", "Solomon Islands", "Somolia", "South Africa", "Sri Lanka", "Swaziland", "Tanzania", "Tonga", "Trinidad and Tobago", "Tuvalu", "Uganda", "United Kingdom", "United States", "Vanuatu", "Zambia", "Zimbabwe"],
              "Arabic": ["Algeria", "Bahrain", "Chad", "Comoros", "Djibouti", "Egypt", "Eritrea", "Iraq", "Israel", "Jordan", "Kuwait", "Lebanon", "Libya", "Morocco", "Niger", "Oman", "Palestinian Territories", "Quatar", "Saudi Arabia", "Somalia", "Sudan", "Syria", "Tunisia", "United Arab Emirates", "Western Sahara", "Yemen", "Mauritania", "Senegal"],
              "Portuguese": ["Angola", "Brazil", "Cape Verde", "East Timor", "Guinea-Bissau", "Macau", "Mozambique", "Portugal"],
              "Russian": ["Belarus", "Kazakhstan", "Kyyrgyzstan", "Russia"],
              "Japanese": ["Japan", "Palau"],
              "German": ["Austria", "Belgium", "Germany", "Italy", "Liechtenstein", "Luxembourg", "Poland", "Siwtzerland"],
              "Korean": ["North Korea", "South Korea"],
              "Vietnamese": ["Vietnam"],
              "French": ["Belgium", "Benin", "Burkina Faso", "Burundi", "Cameroon", "Canada", "Central African Republic", "Chad", "Comoros", "Congo-Brazzaville", "Congo-Kinshasa", "Côte d'Ivoire", "Djibouti", "Equatorial Guinea", "France", "French Polynesia", "Gabon", "Guernsey", "Guinea", "Haiti", "Italy", "Jersey", "Lebanon", "Luxembourg", "Madagascar", "Mali", "Martinique", "Mauritius", "Mayotte", "Monaco", "New Caledonia", "Niger", "Rwanda", "Senegal", "Seychelles", "Switzerland", "Togo", "Vanuatu"],
              "Italian": ["Croatia", "Italy", "San Marino", "Slovenia", "Switzerland"],
              "Turkish": ["Bulgaria", "Cyprus", "Turkey"],
              "Polish": ["Poland"],
              "Thai": ["Thailand"]
              },


          total_cntrs: 0,
          total_users: 0,
          app_cntrs: 0,
          app_users: 0,
          percent: 0,

          getLangs: function () {
             var langs = [];
             var items = this.$.langlist.getElementsByTagName('input');

             for (var i=0; i < items.length; i++) {           
                 if (items[i].checked) {                 
                     langs.push(items[i].value);
                 } 
             }
             return langs;
          },   


          getCountries: function (langs) {
            var self = this;
            var cntrs = langs.reduce(function(prev, item) { 
                return prev.concat(self.langs[item]);
            }, []);

            cntrs = cntrs.filter(function(item, pos) {
                return cntrs.indexOf(item) == pos;
            });

            return this.countries.filter(function(item, pos) {
                return cntrs.indexOf(item[1]) >= 0;
            });
          },

          usersCount: function (items) {
              if (items.length === 0) {
                return 0;
              }
              var result = items.reduce(function(sum, item) {
                  return sum + item[2];
              }, 0);

              return parseInt((result/(3*Math.pow(10, 6))).toFixed(1));
          },


          redraw: function (e) { 

             var self = this;        

             var data = [['Code', 'Country',   'iOS Devices']];
                           
             var langs = this.getLangs();

             var app_cntrs  = this.getCountries(langs);
             this.app_users = this.usersCount(app_cntrs);
             this.app_cntrs = app_cntrs.length;

             
             this.total_users = this.usersCount(this.countries);
             this.total_cntrs = this.countries.length;


             console.log("Total users:", this.total_users, "App users:", this.app_users);

        
             for (i = 0; i < app_cntrs.length; i++) {
               data.push(app_cntrs[i]);
             }

             if (data.length > 1) {  
                this.$.geo.data = data;                   
             }
             else {
               this.$.geo.data = [['Code', 'Country',   'iOS Devices'],
                           ['KI', 'Kiribati', 0]];
             }
             //this.$.geo.drawChart();      


            var rows = this.langs_keys.map(function(item) {


             var cntrs  = self.getCountries([item]);
             var users = self.usersCount(cntrs);
              
                return [item, users];
            });

            rows = rows.sort(function (a,b) {

                if (a[1] > b[1]) {
                  return 1;
                }
                if (a[1] < b[1]) {
                  return -1;
                }
                return 0;
            });

            this.$.pie.rows = rows;

            this.percent = parseInt(this.app_users/(this.total_users/100));

             //this.$.pie.rows=[["Jan", 31],["Feb", 28],["Mar", 31],["Apr", 30],["May", 31],["Jun", 30]];
             //this.$.pie.drawChart();  
         },




          domReady: function () {

     
            

            this.langs_keys = Object.keys(App.langs);



              this.$.geo.options = {
                //region: '002', // Africa
                colorAxis: {colors: ['#00853f','#00853f']},
                //backgroundColor: '#81d4fa',
                backgroundColor: '#ffffff',
                datalessRegionColor: '#eee',
                defaultColor: '#f5f5f5',
              };
              
            //this.$.pie.options = {"title": "Distribution of days in 2001H1"};

            this.$.pie.options = {
               // colors: ['red','#004411']
              }

            
            this.$.pie.cols=[{"label": "Language", "type": "string"},{"label": "Uers", "type": "number"}];
            
        }

     };

      // Load the Visualization API and the piechart package.
      google.load('visualization', '1', {'packages':['piechart', 'geochart']});
      google.setOnLoadCallback(drawRegionsMap);

      // Set a callback to run when the Google Visualization API is loaded.
      function drawRegionsMap() {

        var data = google.visualization.arrayToDataTable([
          ['Country', 'Popularity'],
          ['Germany', 200],
          ['United States', 300],
          ['Brazil', 400],
          ['Canada', 500],
          ['France', 600],
          ['RU', 700]
        ]);

        var options = {};

        var chart = new google.visualization.GeoChart(document.getElementById('regions_div'));

        chart.draw(data, options);
      }