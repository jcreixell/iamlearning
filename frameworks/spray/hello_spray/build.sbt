val akkaV = "2.3.9"
val sprayV = "1.3.3"
 
libraryDependencies ++= Seq(
    "io.spray" %% "spray-can" % sprayV,
    "io.spray" %% "spray-routing" % sprayV,
    "com.typesafe.akka" %% "akka-actor" % akkaV
  )
