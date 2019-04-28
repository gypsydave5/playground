package main

import "html/template"

func viewer() *template.Template {
	return template.Must(template.New("viewer").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width" />
    <title>MandlebrotZoomer</title>
  </head>
  <body>
    <img src='/mandlebrot?height=512&width=512'
         id="mandelbrot"></img>
  </body>
  <script type="application/javascript" id="">
    const baseURL = new URL('/mandlebrot', window.location.protocol + '//' + window.location.host)
    const settings = {
      height: 512,
      width: 512,
      colour: true,
      x: 0.0,
      y: 0.0,
      zoom: 1.0,
      bounds: 2,
      iterations: 20
    }

    const m = document.getElementById('mandelbrot')

    m.addEventListener('click', e => {
      const [x, y] = calculateCoordinates(e, settings)
      settings.x = settings.x + x
      settings.y = settings.y + y
      settings.zoom = settings.zoom / 2
      settings.iterations = newIteration(settings)

      const qs = new URLSearchParams(settings)
      e.target.src = baseURL.toString() + "?" + qs.toString()
    })

function newIteration (settings) {
  if (settings.iterations === 255) {
    return 255
  }
  if ((settings.iterations / settings.zoom) > 255) {
    return 255
  }
  return settings.iterations / settings.zoom
}

function calculateCoordinates (e, settings) {
  const rect = e.target.getBoundingClientRect()
  const x = e.clientX - rect.left;
  const y = e.clientY - rect.top;

  const xmid = settings.height / 2
  const nextBounds = settings.bounds * settings.zoom

  const newX = (x - xmid) / (xmid / nextBounds)
  const newY = (y - xmid) / (xmid / nextBounds)
  return [newX, newY]
}
  </script>
</html>
  `))
}
