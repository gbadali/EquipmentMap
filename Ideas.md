I am trying to make the breadcrumbs at the top of the page.
in order to do that I need to have the location back to the top level of the graph.

    <nav aria-label="breadcrumb">
    	<ul>
    		<li><a href="#">Home</a></li>
    		<li><a href="#">Services</a></li>
    		<li>Design</li>
    	</ul>
    </nav>

So go gives templ some object and then templ itereates over that object creating `<li>` objects down the graph. What should i make the "thing" that gets passed to templ? An array? `[ship, gen1, ssdg1, fuel pump]`? Where should I put it, i think the easiest place to put it would be in the handlers? its probably not the "right" place to put it but it is probably the easiest.
