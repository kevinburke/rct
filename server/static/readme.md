offline jquery:

You can achieve it like this:

<script src="//ajax.googleapis.com/ajax/libs/jquery/1.2.6/jquery.min.js"></script>
<script>if (!window.jQuery) { document.write('<script src="/path/to/your/jquery"><\/script>'); }
</script>

<script src="//ajax.googleapis.com/ajax/libs/jquery/1.7.2/jquery.min.js"></script>
<script>window.jQuery || document.write('<script src="path/to/your/jquery"><\/script>')</script>

This should be in your page's <head> and any jQuery ready event handlers should be in the <body> to avoid errors (although it's not fool-proof!).

One more reason to not use Google-hosted jQuery is that in some countries, Google's domain name is banned.
