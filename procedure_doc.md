## UI
#### Classes
```
Button
<doc>
<summary>
A text label. See <see cref="M:UI.Panel.AddButton" />.
</summary>
</doc>
```
```
Canvas
<doc>
<summary>
A canvas for user interface elements. See <see cref="M:UI.StockCanvas" /> and <see cref="M:UI.AddCanvas" />.
</summary>
</doc>
```
```
InputField
<doc>
<summary>
An input field. See <see cref="M:UI.Panel.AddInputField" />.
</summary>
</doc>
```
```
Panel
<doc>
<summary>
A container for user interface elements. See <see cref="M:UI.Canvas.AddPanel" />.
</summary>
</doc>
```
```
RectTransform
<doc>
<summary>
A Unity engine Rect Transform for a UI object.
See the <a href="https://docs.unity3d.com/Manual/class-RectTransform.html">Unity manual</a> for more details.
</summary>
</doc>
```
```
Text
<doc>
<summary>
A text label. See <see cref="M:UI.Panel.AddText" />.
</summary>
</doc>
```
### AddCanvas

```
<doc>
<summary>
Add a new canvas.
</summary>
<remarks>
If you want to add UI elements to KSPs stock UI canvas, use <see cref="M:UI.StockCanvas" />.
</remarks>
</doc>
```
#### Parameters 
#### Return 
```
Canvas```
### Message

```
<doc>
<summary>
Display a message on the screen.
</summary>
<remarks>
The message appears just like a stock message, for example quicksave or quickload messages.
</remarks>
<param name="content">Message content.</param>
<param name="duration">Duration before the message disappears, in seconds.</param>
<param name="position">Position to display the message.</param>
</doc>
```
#### Parameters 
```
content		
```
```
duration		
```
```
position		MessagePosition
```
#### Return 
```
```
### Clear

```
<doc>
<summary>
Remove all user interface elements.
</summary>
<param name="clientOnly">If true, only remove objects created by the calling client.</param>
</doc>
```
#### Parameters 
```
clientOnly		
```
#### Return 
```
```
### get_StockCanvas

```
<doc>
<summary>
The stock UI canvas.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Canvas```
### Button_Remove

```
<doc>
<summary>
Remove the UI object.
</summary>
</doc>
```
#### Parameters 
```
this		Button
```
#### Return 
```
```
### Button_get_RectTransform

```
<doc>
<summary>
The rect transform for the text.
</summary>
</doc>
```
#### Parameters 
```
this		Button
```
#### Return 
```
RectTransform```
### Button_get_Text

```
<doc>
<summary>
The text for the button.
</summary>
</doc>
```
#### Parameters 
```
this		Button
```
#### Return 
```
Text```
### Button_get_Clicked

```
<doc>
<summary>
Whether the button has been clicked.
</summary>
<remarks>
This property is set to true when the user clicks the button.
A client script should reset the property to false in order to detect subsequent button presses.
</remarks>
</doc>
```
#### Parameters 
```
this		Button
```
#### Return 
```
```
### Button_set_Clicked

```
<doc>
<summary>
Whether the button has been clicked.
</summary>
<remarks>
This property is set to true when the user clicks the button.
A client script should reset the property to false in order to detect subsequent button presses.
</remarks>
</doc>
```
#### Parameters 
```
this		Button
```
```
value		
```
#### Return 
```
```
### Button_get_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Button
```
#### Return 
```
```
### Button_set_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Button
```
```
value		
```
#### Return 
```
```
### Canvas_AddPanel

```
<doc>
<summary>
Create a new container for user interface elements.
</summary>
<param name="visible">Whether the panel is visible.</param>
</doc>
```
#### Parameters 
```
this		Canvas
```
```
visible		
```
#### Return 
```
Panel```
### Canvas_AddText

```
<doc>
<summary>
Add text to the canvas.
</summary>
<param name="content">The text.</param>
<param name="visible">Whether the text is visible.</param>
</doc>
```
#### Parameters 
```
this		Canvas
```
```
content		
```
```
visible		
```
#### Return 
```
Text```
### Canvas_AddInputField

```
<doc>
<summary>
Add an input field to the canvas.
</summary>
<param name="visible">Whether the input field is visible.</param>
</doc>
```
#### Parameters 
```
this		Canvas
```
```
visible		
```
#### Return 
```
InputField```
### Canvas_AddButton

```
<doc>
<summary>
Add a button to the canvas.
</summary>
<param name="content">The label for the button.</param>
<param name="visible">Whether the button is visible.</param>
</doc>
```
#### Parameters 
```
this		Canvas
```
```
content		
```
```
visible		
```
#### Return 
```
Button```
### Canvas_Remove

```
<doc>
<summary>
Remove the UI object.
</summary>
</doc>
```
#### Parameters 
```
this		Canvas
```
#### Return 
```
```
### Canvas_get_RectTransform

```
<doc>
<summary>
The rect transform for the canvas.
</summary>
</doc>
```
#### Parameters 
```
this		Canvas
```
#### Return 
```
RectTransform```
### Canvas_get_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Canvas
```
#### Return 
```
```
### Canvas_set_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Canvas
```
```
value		
```
#### Return 
```
```
### InputField_Remove

```
<doc>
<summary>
Remove the UI object.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
```
### InputField_get_RectTransform

```
<doc>
<summary>
The rect transform for the input field.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
RectTransform```
### InputField_get_Value

```
<doc>
<summary>
The value of the input field.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
```
### InputField_set_Value

```
<doc>
<summary>
The value of the input field.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
```
value		
```
#### Return 
```
```
### InputField_get_Text

```
<doc>
<summary>
The text component of the input field.
</summary>
<remarks>
Use <see cref="M:UI.InputField.Value" /> to get and set the value in the field.
This object can be used to alter the style of the input field's text.
</remarks>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
Text```
### InputField_get_Changed

```
<doc>
<summary>
Whether the input field has been changed.
</summary>
<remarks>
This property is set to true when the user modifies the value of the input field.
A client script should reset the property to false in order to detect subsequent changes.
</remarks>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
```
### InputField_set_Changed

```
<doc>
<summary>
Whether the input field has been changed.
</summary>
<remarks>
This property is set to true when the user modifies the value of the input field.
A client script should reset the property to false in order to detect subsequent changes.
</remarks>
</doc>
```
#### Parameters 
```
this		InputField
```
```
value		
```
#### Return 
```
```
### InputField_get_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
#### Return 
```
```
### InputField_set_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		InputField
```
```
value		
```
#### Return 
```
```
### Panel_AddPanel

```
<doc>
<summary>
Create a panel within this panel.
</summary>
<param name="visible">Whether the new panel is visible.</param>
</doc>
```
#### Parameters 
```
this		Panel
```
```
visible		
```
#### Return 
```
Panel```
### Panel_AddText

```
<doc>
<summary>
Add text to the panel.
</summary>
<param name="content">The text.</param>
<param name="visible">Whether the text is visible.</param>
</doc>
```
#### Parameters 
```
this		Panel
```
```
content		
```
```
visible		
```
#### Return 
```
Text```
### Panel_AddInputField

```
<doc>
<summary>
Add an input field to the panel.
</summary>
<param name="visible">Whether the input field is visible.</param>
</doc>
```
#### Parameters 
```
this		Panel
```
```
visible		
```
#### Return 
```
InputField```
### Panel_AddButton

```
<doc>
<summary>
Add a button to the panel.
</summary>
<param name="content">The label for the button.</param>
<param name="visible">Whether the button is visible.</param>
</doc>
```
#### Parameters 
```
this		Panel
```
```
content		
```
```
visible		
```
#### Return 
```
Button```
### Panel_Remove

```
<doc>
<summary>
Remove the UI object.
</summary>
</doc>
```
#### Parameters 
```
this		Panel
```
#### Return 
```
```
### Panel_get_RectTransform

```
<doc>
<summary>
The rect transform for the panel.
</summary>
</doc>
```
#### Parameters 
```
this		Panel
```
#### Return 
```
RectTransform```
### Panel_get_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Panel
```
#### Return 
```
```
### Panel_set_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Panel
```
```
value		
```
#### Return 
```
```
### RectTransform_get_Position

```
<doc>
<summary>
Position of the rectangles pivot point relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_Position

```
<doc>
<summary>
Position of the rectangles pivot point relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_LocalPosition

```
<doc>
<summary>
Position of the rectangles pivot point relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_LocalPosition

```
<doc>
<summary>
Position of the rectangles pivot point relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_Size

```
<doc>
<summary>
Width and height of the rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_Size

```
<doc>
<summary>
Width and height of the rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_UpperRight

```
<doc>
<summary>
Position of the rectangles upper right corner relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_UpperRight

```
<doc>
<summary>
Position of the rectangles upper right corner relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_LowerLeft

```
<doc>
<summary>
Position of the rectangles lower left corner relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_LowerLeft

```
<doc>
<summary>
Position of the rectangles lower left corner relative to the anchors.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_set_Anchor

```
<doc>
<summary>
Set the minimum and maximum anchor points as a fraction of the size of the parent rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_AnchorMax

```
<doc>
<summary>
The anchor point for the lower left corner of the rectangle defined as a fraction of the size of the parent rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_AnchorMax

```
<doc>
<summary>
The anchor point for the lower left corner of the rectangle defined as a fraction of the size of the parent rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_AnchorMin

```
<doc>
<summary>
The anchor point for the upper right corner of the rectangle defined as a fraction of the size of the parent rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_AnchorMin

```
<doc>
<summary>
The anchor point for the upper right corner of the rectangle defined as a fraction of the size of the parent rectangle.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_Pivot

```
<doc>
<summary>
Location of the pivot point around which the rectangle rotates, defined as a fraction of the size of the rectangle itself.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_Pivot

```
<doc>
<summary>
Location of the pivot point around which the rectangle rotates, defined as a fraction of the size of the rectangle itself.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_Rotation

```
<doc>
<summary>
Rotation, as a quaternion, of the object around its pivot point.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_Rotation

```
<doc>
<summary>
Rotation, as a quaternion, of the object around its pivot point.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### RectTransform_get_Scale

```
<doc>
<summary>
Scale factor applied to the object in the x, y and z dimensions.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
#### Return 
```
```
### RectTransform_set_Scale

```
<doc>
<summary>
Scale factor applied to the object in the x, y and z dimensions.
</summary>
</doc>
```
#### Parameters 
```
this		RectTransform
```
```
value		
```
#### Return 
```
```
### Text_Remove

```
<doc>
<summary>
Remove the UI object.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_get_RectTransform

```
<doc>
<summary>
The rect transform for the text.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
RectTransform```
### Text_get_AvailableFonts

```
<doc>
<summary>
A list of all available fonts.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_get_Content

```
<doc>
<summary>
The text string
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Content

```
<doc>
<summary>
The text string
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Font

```
<doc>
<summary>
Name of the font
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Font

```
<doc>
<summary>
Name of the font
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Size

```
<doc>
<summary>
Font size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Size

```
<doc>
<summary>
Font size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Style

```
<doc>
<summary>
Font style.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
FontStyle```
### Text_set_Style

```
<doc>
<summary>
Font style.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		FontStyle
```
#### Return 
```
```
### Text_get_Alignment

```
<doc>
<summary>
Alignment.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
TextAnchor```
### Text_set_Alignment

```
<doc>
<summary>
Alignment.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		TextAnchor
```
#### Return 
```
```
### Text_get_LineSpacing

```
<doc>
<summary>
Line spacing.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_LineSpacing

```
<doc>
<summary>
Line spacing.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Visible

```
<doc>
<summary>
Whether the UI object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
## RemoteTech
#### Classes
```
Antenna
<doc>
<summary>
A RemoteTech antenna. Obtained by calling <see cref="M:RemoteTech.Comms.Antennas" /> or <see cref="M:RemoteTech.Antenna" />.
</summary>
</doc>
```
```
Comms
<doc>
<summary>
Communications for a vessel.
</summary>
</doc>
```
### Comms

```
<doc>
<summary>
Get a communications object, representing the communication capability of a particular vessel.
</summary>
</doc>
```
#### Parameters 
```
vessel		Vessel
```
#### Return 
```
Comms```
### Antenna

```
<doc>
<summary>
Get the antenna object for a particular part.
</summary>
</doc>
```
#### Parameters 
```
part		Part
```
#### Return 
```
Antenna```
### get_Available

```
<doc>
<summary>
Whether RemoteTech is installed.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_GroundStations

```
<doc>
<summary>
The names of the ground stations.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### Antenna_get_Part

```
<doc>
<summary>
Get the part containing this antenna.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
Part```
### Antenna_get_HasConnection

```
<doc>
<summary>
Whether the antenna has a connection.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_Target

```
<doc>
<summary>
The object that the antenna is targetting.
This property can be used to set the target to <see cref="M:RemoteTech.Target.None" /> or <see cref="M:RemoteTech.Target.ActiveVessel" />.
To set the target to a celestial body, ground station or vessel see <see cref="M:RemoteTech.Antenna.TargetBody" />,
<see cref="M:RemoteTech.Antenna.TargetGroundStation" /> and <see cref="M:RemoteTech.Antenna.TargetVessel" />.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
Target```
### Antenna_set_Target

```
<doc>
<summary>
The object that the antenna is targetting.
This property can be used to set the target to <see cref="M:RemoteTech.Target.None" /> or <see cref="M:RemoteTech.Target.ActiveVessel" />.
To set the target to a celestial body, ground station or vessel see <see cref="M:RemoteTech.Antenna.TargetBody" />,
<see cref="M:RemoteTech.Antenna.TargetGroundStation" /> and <see cref="M:RemoteTech.Antenna.TargetVessel" />.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		Target
```
#### Return 
```
```
### Antenna_get_TargetBody

```
<doc>
<summary>
The celestial body the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
CelestialBody```
### Antenna_set_TargetBody

```
<doc>
<summary>
The celestial body the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		CelestialBody
```
#### Return 
```
```
### Antenna_get_TargetGroundStation

```
<doc>
<summary>
The ground station the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_set_TargetGroundStation

```
<doc>
<summary>
The ground station the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		
```
#### Return 
```
```
### Antenna_get_TargetVessel

```
<doc>
<summary>
The vessel the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
Vessel```
### Antenna_set_TargetVessel

```
<doc>
<summary>
The vessel the antenna is targetting.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		Vessel
```
#### Return 
```
```
### Comms_SignalDelayToVessel

```
<doc>
<summary>
The signal delay between the this vessel and another vessel, in seconds.
</summary>
<param name="other"></param>
</doc>
```
#### Parameters 
```
this		Comms
```
```
other		Vessel
```
#### Return 
```
```
### Comms_get_Vessel

```
<doc>
<summary>
Get the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
Vessel```
### Comms_get_HasLocalControl

```
<doc>
<summary>
Whether the vessel can be controlled locally.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_HasFlightComputer

```
<doc>
<summary>
Whether the vessel has a flight computer on board.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_HasConnection

```
<doc>
<summary>
Whether the vessel has any connection.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_HasConnectionToGroundStation

```
<doc>
<summary>
Whether the vessel has a connection to a ground station.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_SignalDelay

```
<doc>
<summary>
The shortest signal delay to the vessel, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_SignalDelayToGroundStation

```
<doc>
<summary>
The signal delay between the vessel and the closest ground station, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_Antennas

```
<doc>
<summary>
The antennas for this vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
## KerbalAlarmClock
#### Classes
```
Alarm
<doc>
<summary>
Represents an alarm. Obtained by calling
<see cref="M:KerbalAlarmClock.Alarms" />,
<see cref="M:KerbalAlarmClock.AlarmWithName" /> or
<see cref="M:KerbalAlarmClock.AlarmsWithType" />.
</summary>
</doc>
```
### AlarmWithName

```
<doc>
<summary>
Get the alarm with the given <paramref name="name" />, or <c>null</c>
if no alarms have that name. If more than one alarm has the name,
only returns one of them.
</summary>
<param name="name">Name of the alarm to search for.</param>
</doc>
```
#### Parameters 
```
name		
```
#### Return 
```
Alarm```
### AlarmsWithType

```
<doc>
<summary>
Get a list of alarms of the specified <paramref name="type" />.
</summary>
<param name="type">Type of alarm to return.</param>
</doc>
```
#### Parameters 
```
type		AlarmType
```
#### Return 
```
```
### CreateAlarm

```
<doc>
<summary>
Create a new alarm and return it.
</summary>
<param name="type">Type of the new alarm.</param>
<param name="name">Name of the new alarm.</param>
<param name="ut">Time at which the new alarm should trigger.</param>
</doc>
```
#### Parameters 
```
type		AlarmType
```
```
name		
```
```
ut		
```
#### Return 
```
Alarm```
### get_Available

```
<doc>
<summary>
Whether Kerbal Alarm Clock is available.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_Alarms

```
<doc>
<summary>
A list of all the alarms.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### Alarm_Remove

```
<doc>
<summary>
Removes the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_get_Action

```
<doc>
<summary>
The action that the alarm triggers.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
AlarmAction```
### Alarm_set_Action

```
<doc>
<summary>
The action that the alarm triggers.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		AlarmAction
```
#### Return 
```
```
### Alarm_get_Margin

```
<doc>
<summary>
The number of seconds before the event that the alarm will fire.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_Margin

```
<doc>
<summary>
The number of seconds before the event that the alarm will fire.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_Time

```
<doc>
<summary>
The time at which the alarm will fire.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_Time

```
<doc>
<summary>
The time at which the alarm will fire.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_Type

```
<doc>
<summary>
The type of the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
AlarmType```
### Alarm_get_ID

```
<doc>
<summary>
The unique identifier for the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_get_Name

```
<doc>
<summary>
The short name of the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_Name

```
<doc>
<summary>
The short name of the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_Notes

```
<doc>
<summary>
The long description of the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_Notes

```
<doc>
<summary>
The long description of the alarm.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_Remaining

```
<doc>
<summary>
The number of seconds until the alarm will fire.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_get_Repeat

```
<doc>
<summary>
Whether the alarm will be repeated after it has fired.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_Repeat

```
<doc>
<summary>
Whether the alarm will be repeated after it has fired.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_RepeatPeriod

```
<doc>
<summary>
The time delay to automatically create an alarm after it has fired.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
```
### Alarm_set_RepeatPeriod

```
<doc>
<summary>
The time delay to automatically create an alarm after it has fired.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		
```
#### Return 
```
```
### Alarm_get_Vessel

```
<doc>
<summary>
The vessel that the alarm is attached to.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
Vessel```
### Alarm_set_Vessel

```
<doc>
<summary>
The vessel that the alarm is attached to.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		Vessel
```
#### Return 
```
```
### Alarm_get_XferOriginBody

```
<doc>
<summary>
The celestial body the vessel is departing from.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
CelestialBody```
### Alarm_set_XferOriginBody

```
<doc>
<summary>
The celestial body the vessel is departing from.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		CelestialBody
```
#### Return 
```
```
### Alarm_get_XferTargetBody

```
<doc>
<summary>
The celestial body the vessel is arriving at.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
#### Return 
```
CelestialBody```
### Alarm_set_XferTargetBody

```
<doc>
<summary>
The celestial body the vessel is arriving at.
</summary>
</doc>
```
#### Parameters 
```
this		Alarm
```
```
value		CelestialBody
```
#### Return 
```
```
## InfernalRobotics
#### Classes
```
Servo
<doc>
<summary>
Represents a servo. Obtained using
<see cref="M:InfernalRobotics.ServoGroup.Servos" />,
<see cref="M:InfernalRobotics.ServoGroup.ServoWithName" />
or <see cref="M:InfernalRobotics.ServoWithName" />.
</summary>
</doc>
```
```
ServoGroup
<doc>
<summary>
A group of servos, obtained by calling <see cref="M:InfernalRobotics.ServoGroups" />
or <see cref="M:InfernalRobotics.ServoGroupWithName" />. Represents the "Servo Groups"
in the InfernalRobotics UI.
</summary>
</doc>
```
### ServoGroups

```
<doc>
<summary>
A list of all the servo groups in the given <paramref name="vessel" />.
</summary>
</doc>
```
#### Parameters 
```
vessel		Vessel
```
#### Return 
```
```
### ServoGroupWithName

```
<doc>
<summary>
Returns the servo group in the given <paramref name="vessel" /> with the given <paramref name="name" />,
or <c>null</c> if none exists. If multiple servo groups have the same name, only one of them is returned.
</summary>
<param name="vessel">Vessel to check.</param>
<param name="name">Name of servo group to find.</param>
</doc>
```
#### Parameters 
```
vessel		Vessel
```
```
name		
```
#### Return 
```
ServoGroup```
### ServoWithName

```
<doc>
<summary>
Returns the servo in the given <paramref name="vessel" /> with the given <paramref name="name" /> or
<c>null</c> if none exists. If multiple servos have the same name, only one of them is returned.
</summary>
<param name="vessel">Vessel to check.</param>
<param name="name">Name of the servo to find.</param>
</doc>
```
#### Parameters 
```
vessel		Vessel
```
```
name		
```
#### Return 
```
Servo```
### get_Available

```
<doc>
<summary>
Whether Infernal Robotics is installed.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### Servo_MoveRight

```
<doc>
<summary>
Moves the servo to the right.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_MoveLeft

```
<doc>
<summary>
Moves the servo to the left.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_MoveCenter

```
<doc>
<summary>
Moves the servo to the center.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_MoveNextPreset

```
<doc>
<summary>
Moves the servo to the next preset.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_MovePrevPreset

```
<doc>
<summary>
Moves the servo to the previous preset.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_MoveTo

```
<doc>
<summary>
Moves the servo to <paramref name="position" /> and sets the
speed multiplier to <paramref name="speed" />.
</summary>
<param name="position">The position to move the servo to.</param>
<param name="speed">Speed multiplier for the movement.</param>
</doc>
```
#### Parameters 
```
this		Servo
```
```
position		
```
```
speed		
```
#### Return 
```
```
### Servo_Stop

```
<doc>
<summary>
Stops the servo.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_Name

```
<doc>
<summary>
The name of the servo.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_Name

```
<doc>
<summary>
The name of the servo.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_Part

```
<doc>
<summary>
The part containing the servo.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
Part```
### Servo_set_Highlight

```
<doc>
<summary>
Whether the servo should be highlighted in-game.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_Position

```
<doc>
<summary>
The position of the servo.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_MinConfigPosition

```
<doc>
<summary>
The minimum position of the servo, specified by the part configuration.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_MaxConfigPosition

```
<doc>
<summary>
The maximum position of the servo, specified by the part configuration.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_MinPosition

```
<doc>
<summary>
The minimum position of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_MinPosition

```
<doc>
<summary>
The minimum position of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_MaxPosition

```
<doc>
<summary>
The maximum position of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_MaxPosition

```
<doc>
<summary>
The maximum position of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_ConfigSpeed

```
<doc>
<summary>
The speed multiplier of the servo, specified by the part configuration.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_Speed

```
<doc>
<summary>
The speed multiplier of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_Speed

```
<doc>
<summary>
The speed multiplier of the servo, specified by the in-game tweak menu.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_CurrentSpeed

```
<doc>
<summary>
The current speed at which the servo is moving.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_CurrentSpeed

```
<doc>
<summary>
The current speed at which the servo is moving.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_Acceleration

```
<doc>
<summary>
The current speed multiplier set in the UI.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_Acceleration

```
<doc>
<summary>
The current speed multiplier set in the UI.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_IsMoving

```
<doc>
<summary>
Whether the servo is moving.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_IsFreeMoving

```
<doc>
<summary>
Whether the servo is freely moving.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_get_IsLocked

```
<doc>
<summary>
Whether the servo is locked.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_IsLocked

```
<doc>
<summary>
Whether the servo is locked.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### Servo_get_IsAxisInverted

```
<doc>
<summary>
Whether the servos axis is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
#### Return 
```
```
### Servo_set_IsAxisInverted

```
<doc>
<summary>
Whether the servos axis is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Servo
```
```
value		
```
#### Return 
```
```
### ServoGroup_ServoWithName

```
<doc>
<summary>
Returns the servo with the given <paramref name="name" /> from this group,
or <c>null</c> if none exists.
</summary>
<param name="name">Name of servo to find.</param>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
name		
```
#### Return 
```
Servo```
### ServoGroup_MoveRight

```
<doc>
<summary>
Moves all of the servos in the group to the right.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_MoveLeft

```
<doc>
<summary>
Moves all of the servos in the group to the left.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_MoveCenter

```
<doc>
<summary>
Moves all of the servos in the group to the center.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_MoveNextPreset

```
<doc>
<summary>
Moves all of the servos in the group to the next preset.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_MovePrevPreset

```
<doc>
<summary>
Moves all of the servos in the group to the previous preset.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_Stop

```
<doc>
<summary>
Stops the servos in the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_get_Name

```
<doc>
<summary>
The name of the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_set_Name

```
<doc>
<summary>
The name of the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
value		
```
#### Return 
```
```
### ServoGroup_get_ForwardKey

```
<doc>
<summary>
The key assigned to be the "forward" key for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_set_ForwardKey

```
<doc>
<summary>
The key assigned to be the "forward" key for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
value		
```
#### Return 
```
```
### ServoGroup_get_ReverseKey

```
<doc>
<summary>
The key assigned to be the "reverse" key for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_set_ReverseKey

```
<doc>
<summary>
The key assigned to be the "reverse" key for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
value		
```
#### Return 
```
```
### ServoGroup_get_Speed

```
<doc>
<summary>
The speed multiplier for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_set_Speed

```
<doc>
<summary>
The speed multiplier for the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
value		
```
#### Return 
```
```
### ServoGroup_get_Expanded

```
<doc>
<summary>
Whether the group is expanded in the InfernalRobotics UI.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_set_Expanded

```
<doc>
<summary>
Whether the group is expanded in the InfernalRobotics UI.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
```
value		
```
#### Return 
```
```
### ServoGroup_get_Servos

```
<doc>
<summary>
The servos that are in the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
### ServoGroup_get_Parts

```
<doc>
<summary>
The parts containing the servos in the group.
</summary>
</doc>
```
#### Parameters 
```
this		ServoGroup
```
#### Return 
```
```
## SpaceCenter
#### Classes
```
AutoPilot
<doc>
<summary>
Provides basic auto-piloting utilities for a vessel.
Created by calling <see cref="M:SpaceCenter.Vessel.AutoPilot" />.
</summary>
<remarks>
If a client engages the auto-pilot and then closes its connection to the server,
the auto-pilot will be disengaged and its target reference frame, direction and roll
reset to default.
</remarks>
</doc>
```
```
Camera
<doc>
<summary>
Controls the game's camera.
Obtained by calling <see cref="M:SpaceCenter.Camera" />.
</summary>
</doc>
```
```
CelestialBody
<doc>
<summary>
Represents a celestial body (such as a planet or moon).
See <see cref="M:SpaceCenter.Bodies" />.
</summary>
</doc>
```
```
CommLink
<doc>
<summary>
Represents a communication node in the network. For example, a vessel or the KSC.
</summary>
</doc>
```
```
CommNode
<doc>
<summary>
Represents a communication node in the network. For example, a vessel or the KSC.
</summary>
</doc>
```
```
Comms
<doc>
<summary>
Used to interact with CommNet for a given vessel.
Obtained by calling <see cref="M:SpaceCenter.Vessel.Comms" />.
</summary>
</doc>
```
```
Contract
<doc>
<summary>
A contract. Can be accessed using <see cref="M:SpaceCenter.ContractManager" />.
</summary>
</doc>
```
```
ContractManager
<doc>
<summary>
Contracts manager.
Obtained by calling <see cref="M:SpaceCenter.ContractManager" />.
</summary>
</doc>
```
```
ContractParameter
<doc>
<summary>
A contract parameter. See <see cref="M:SpaceCenter.Contract.Parameters" />.
</summary>
</doc>
```
```
Control
<doc>
<summary>
Used to manipulate the controls of a vessel. This includes adjusting the
throttle, enabling/disabling systems such as SAS and RCS, or altering the
direction in which the vessel is pointing.
Obtained by calling <see cref="M:SpaceCenter.Vessel.Control" />.
</summary>
<remarks>
Control inputs (such as pitch, yaw and roll) are zeroed when all clients
that have set one or more of these inputs are no longer connected.
</remarks>
</doc>
```
```
CrewMember
<doc>
<summary>
Represents crew in a vessel. Can be obtained using <see cref="M:SpaceCenter.Vessel.Crew" />.
</summary>
</doc>
```
```
Flight
<doc>
<summary>
Used to get flight telemetry for a vessel, by calling <see cref="M:SpaceCenter.Vessel.Flight" />.
All of the information returned by this class is given in the reference frame
passed to that method.
Obtained by calling <see cref="M:SpaceCenter.Vessel.Flight" />.
</summary>
<remarks>
To get orbital information, such as the apoapsis or inclination, see <see cref="T:SpaceCenter.Orbit" />.
</remarks>
</doc>
```
```
Node
<doc>
<summary>
Represents a maneuver node. Can be created using <see cref="M:SpaceCenter.Control.AddNode" />.
</summary>
</doc>
```
```
Orbit
<doc>
<summary>
Describes an orbit. For example, the orbit of a vessel, obtained by calling
<see cref="M:SpaceCenter.Vessel.Orbit" />, or a celestial body, obtained by calling
<see cref="M:SpaceCenter.CelestialBody.Orbit" />.
</summary>
</doc>
```
```
Antenna
<doc>
<summary>
An antenna. Obtained by calling <see cref="M:SpaceCenter.Part.Antenna" />.
</summary>
</doc>
```
```
CargoBay
<doc>
<summary>
A cargo bay. Obtained by calling <see cref="M:SpaceCenter.Part.CargoBay" />.
</summary>
</doc>
```
```
ControlSurface
<doc>
<summary>
An aerodynamic control surface. Obtained by calling <see cref="M:SpaceCenter.Part.ControlSurface" />.
</summary>
</doc>
```
```
Decoupler
<doc>
<summary>
A decoupler. Obtained by calling <see cref="M:SpaceCenter.Part.Decoupler" /></summary>
</doc>
```
```
DockingPort
<doc>
<summary>
A docking port. Obtained by calling <see cref="M:SpaceCenter.Part.DockingPort" /></summary>
</doc>
```
```
Engine
<doc>
<summary>
An engine, including ones of various types.
For example liquid fuelled gimballed engines, solid rocket boosters and jet engines.
Obtained by calling <see cref="M:SpaceCenter.Part.Engine" />.
</summary>
<remarks>
For RCS thrusters <see cref="M:SpaceCenter.Part.RCS" />.
</remarks>
</doc>
```
```
Experiment
<doc>
<summary>
Obtained by calling <see cref="M:SpaceCenter.Part.Experiment" />.
</summary>
</doc>
```
```
Fairing
<doc>
<summary>
A fairing. Obtained by calling <see cref="M:SpaceCenter.Part.Fairing" />.
</summary>
</doc>
```
```
Force
<doc>
<summary>
Obtained by calling <see cref="M:SpaceCenter.Part.AddForce" />.
</summary>
</doc>
```
```
Intake
<doc>
<summary>
An air intake. Obtained by calling <see cref="M:SpaceCenter.Part.Intake" />.
</summary>
</doc>
```
```
LaunchClamp
<doc>
<summary>
A launch clamp. Obtained by calling <see cref="M:SpaceCenter.Part.LaunchClamp" />.
</summary>
</doc>
```
```
Leg
<doc>
<summary>
A landing leg. Obtained by calling <see cref="M:SpaceCenter.Part.Leg" />.
</summary>
</doc>
```
```
Light
<doc>
<summary>
A light. Obtained by calling <see cref="M:SpaceCenter.Part.Light" />.
</summary>
</doc>
```
```
Module
<doc>
<summary>
This can be used to interact with a specific part module. This includes part modules in
stock KSP, and those added by mods.

In KSP, each part has zero or more
<a href="https://wiki.kerbalspaceprogram.com/wiki/CFG_File_Documentation#MODULES">PartModules</a>
associated with it. Each one contains some of the functionality of the part.
For example, an engine has a "ModuleEngines" part module that contains all the
functionality of an engine.
</summary>
</doc>
```
```
Parachute
<doc>
<summary>
A parachute. Obtained by calling <see cref="M:SpaceCenter.Part.Parachute" />.
</summary>
</doc>
```
```
Part
<doc>
<summary>
Represents an individual part. Vessels are made up of multiple parts.
Instances of this class can be obtained by several methods in <see cref="T:SpaceCenter.Parts" />.
</summary>
</doc>
```
```
Parts
<doc>
<summary>
Instances of this class are used to interact with the parts of a vessel.
An instance can be obtained by calling <see cref="M:SpaceCenter.Vessel.Parts" />.
</summary>
</doc>
```
```
Propellant
<doc>
<summary>
A propellant for an engine. Obtains by calling <see cref="M:SpaceCenter.Engine.Propellants" />.
</summary>
</doc>
```
```
RCS
<doc>
<summary>
An RCS block or thruster. Obtained by calling <see cref="M:SpaceCenter.Part.RCS" />.
</summary>
</doc>
```
```
Radiator
<doc>
<summary>
A radiator. Obtained by calling <see cref="M:SpaceCenter.Part.Radiator" />.
</summary>
</doc>
```
```
ReactionWheel
<doc>
<summary>
A reaction wheel. Obtained by calling <see cref="M:SpaceCenter.Part.ReactionWheel" />.
</summary>
</doc>
```
```
ResourceConverter
<doc>
<summary>
A resource converter. Obtained by calling <see cref="M:SpaceCenter.Part.ResourceConverter" />.
</summary>
</doc>
```
```
ResourceHarvester
<doc>
<summary>
A resource harvester (drill). Obtained by calling <see cref="M:SpaceCenter.Part.ResourceHarvester" />.
</summary>
</doc>
```
```
ScienceData
<doc>
<summary>
Obtained by calling <see cref="M:SpaceCenter.Experiment.Data" />.
</summary>
</doc>
```
```
ScienceSubject
<doc>
<summary>
Obtained by calling <see cref="M:SpaceCenter.Experiment.ScienceSubject" />.
</summary>
</doc>
```
```
Sensor
<doc>
<summary>
A sensor, such as a thermometer. Obtained by calling <see cref="M:SpaceCenter.Part.Sensor" />.
</summary>
</doc>
```
```
SolarPanel
<doc>
<summary>
A solar panel. Obtained by calling <see cref="M:SpaceCenter.Part.SolarPanel" />.
</summary>
</doc>
```
```
Thruster
<doc>
<summary>
The component of an <see cref="T:SpaceCenter.Engine" /> or <see cref="T:SpaceCenter.RCS" /> part that generates thrust.
Can obtained by calling <see cref="M:SpaceCenter.Engine.Thrusters" /> or <see cref="M:SpaceCenter.RCS.Thrusters" />.
</summary>
<remarks>
Engines can consist of multiple thrusters.
For example, the S3 KS-25x4 "Mammoth" has four rocket nozzels, and so consists of
four thrusters.
</remarks>
</doc>
```
```
Wheel
<doc>
<summary>
A wheel. Includes landing gear and rover wheels.
Obtained by calling <see cref="M:SpaceCenter.Part.Wheel" />.
Can be used to control the motors, steering and deployment of wheels, among other things.
</summary>
</doc>
```
```
ReferenceFrame
<doc>
<summary>
Represents a reference frame for positions, rotations and
velocities. Contains:
<list type="bullet"><item><description>The position of the origin.</description></item><item><description>The directions of the x, y and z axes.</description></item><item><description>The linear velocity of the frame.</description></item><item><description>The angular velocity of the frame.</description></item></list></summary>
<remarks>
This class does not contain any properties or methods. It is only
used as a parameter to other functions.
</remarks>
</doc>
```
```
Resource
<doc>
<summary>
An individual resource stored within a part.
Created using methods in the <see cref="T:SpaceCenter.Resources" /> class.
</summary>
</doc>
```
```
ResourceTransfer
<doc>
<summary>
Transfer resources between parts.
</summary>
</doc>
```
```
Resources
<doc>
<summary>
Represents the collection of resources stored in a vessel, stage or part.
Created by calling <see cref="M:SpaceCenter.Vessel.Resources" />,
<see cref="M:SpaceCenter.Vessel.ResourcesInDecoupleStage" /> or
<see cref="M:SpaceCenter.Part.Resources" />.
</summary>
</doc>
```
```
Vessel
<doc>
<summary>
These objects are used to interact with vessels in KSP. This includes getting
orbital and flight data, manipulating control inputs and managing resources.
Created using <see cref="M:SpaceCenter.ActiveVessel" /> or <see cref="M:SpaceCenter.Vessels" />.
</summary>
</doc>
```
```
Waypoint
<doc>
<summary>
Represents a waypoint. Can be created using <see cref="M:SpaceCenter.WaypointManager.AddWaypoint" />.
</summary>
</doc>
```
```
WaypointManager
<doc>
<summary>
Waypoints are the location markers you can see on the map view showing you where contracts are targeted for.
With this structure, you can obtain coordinate data for the locations of these waypoints.
Obtained by calling <see cref="M:SpaceCenter.WaypointManager" />.
</summary>
</doc>
```
### ClearTarget

```
<doc>
<summary>
Clears the current target.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### LaunchableVessels

```
<doc>
<summary>
Returns a list of vessels from the given <paramref name="craftDirectory" />
that can be launched.
</summary>
<param name="craftDirectory">Name of the directory in the current saves
"Ships" directory. For example <c>"VAB"</c> or <c>"SPH"</c>.</param>
</doc>
```
#### Parameters 
```
craftDirectory		
```
#### Return 
```
```
### LaunchVessel

```
<doc>
<summary>
Launch a vessel.
</summary>
<param name="craftDirectory">Name of the directory in the current saves
"Ships" directory, that contains the craft file.
For example <c>"VAB"</c> or <c>"SPH"</c>.</param>
<param name="name">Name of the vessel to launch. This is the name of the ".craft" file
in the save directory, without the ".craft" file extension.</param>
<param name="launchSite">Name of the launch site. For example <c>"LaunchPad"</c> or
<c>"Runway"</c>.</param>
<param name="recover">If true and there is a vessel on the launch site,
recover it before launching.</param>
<remarks>
Throws an exception if any of the games pre-flight checks fail.
</remarks>
</doc>
```
#### Parameters 
```
craftDirectory		
```
```
name		
```
```
launchSite		
```
```
recover		
```
#### Return 
```
```
### LaunchVesselFromVAB

```
<doc>
<summary>
Launch a new vessel from the VAB onto the launchpad.
</summary>
<param name="name">Name of the vessel to launch.</param>
<param name="recover">If true and there is a vessel on the launch pad,
recover it before launching.</param>
<remarks>
This is equivalent to calling <see cref="M:SpaceCenter.LaunchVessel" /> with the craft directory
set to "VAB" and the launch site set to "LaunchPad".
Throws an exception if any of the games pre-flight checks fail.
</remarks>
</doc>
```
#### Parameters 
```
name		
```
```
recover		
```
#### Return 
```
```
### LaunchVesselFromSPH

```
<doc>
<summary>
Launch a new vessel from the SPH onto the runway.
</summary>
<param name="name">Name of the vessel to launch.</param>
<param name="recover">If true and there is a vessel on the runway,
recover it before launching.</param>
<remarks>
This is equivalent to calling <see cref="M:SpaceCenter.LaunchVessel" /> with the craft directory
set to "SPH" and the launch site set to "Runway".
Throws an exception if any of the games pre-flight checks fail.
</remarks>
</doc>
```
#### Parameters 
```
name		
```
```
recover		
```
#### Return 
```
```
### Save

```
<doc>
<summary>
Save the game with a given name.
This will create a save file called <c>name.sfs</c> in the folder of the
current save game.
</summary>
</doc>
```
#### Parameters 
```
name		
```
#### Return 
```
```
### Load

```
<doc>
<summary>
Load the game with the given name.
This will create a load a save file called <c>name.sfs</c> from the folder of the
current save game.
</summary>
</doc>
```
#### Parameters 
```
name		
```
#### Return 
```
```
### Quicksave

```
<doc>
<summary>
Save a quicksave.
</summary>
<remarks>
This is the same as calling <see cref="M:SpaceCenter.Save" /> with the name "quicksave".
</remarks>
</doc>
```
#### Parameters 
#### Return 
```
```
### Quickload

```
<doc>
<summary>
Load a quicksave.
</summary>
<remarks>
This is the same as calling <see cref="M:SpaceCenter.Load" /> with the name "quicksave".
</remarks>
</doc>
```
#### Parameters 
#### Return 
```
```
### CanRailsWarpAt

```
<doc>
<summary>
Returns <c>true</c> if regular "on-rails" time warp can be used, at the specified warp
<paramref name="factor" />. The maximum time warp rate is limited by various things,
including how close the active vessel is to a planet. See
<a href="https://wiki.kerbalspaceprogram.com/wiki/Time_warp">the KSP wiki</a>
for details.
</summary>
<param name="factor">The warp factor to check.</param>
</doc>
```
#### Parameters 
```
factor		
```
#### Return 
```
```
### WarpTo

```
<doc>
<summary>
Uses time acceleration to warp forward to a time in the future, specified
by universal time <paramref name="ut" />. This call blocks until the desired
time is reached. Uses regular "on-rails" or physical time warp as appropriate.
For example, physical time warp is used when the active vessel is traveling
through an atmosphere. When using regular "on-rails" time warp, the warp
rate is limited by <paramref name="maxRailsRate" />, and when using physical
time warp, the warp rate is limited by <paramref name="maxPhysicsRate" />.
</summary>
<param name="ut">The universal time to warp to, in seconds.</param>
<param name="maxRailsRate">The maximum warp rate in regular "on-rails" time warp.
</param>
<param name="maxPhysicsRate">The maximum warp rate in physical time warp.</param>
<returns>When the time warp is complete.</returns>
</doc>
```
#### Parameters 
```
ut		
```
```
maxRailsRate		
```
```
maxPhysicsRate		
```
#### Return 
```
```
### TransformPosition

```
<doc>
<summary>
Converts a position from one reference frame to another.
</summary>
<param name="position">Position, as a vector, in reference frame
<paramref name="from" />.</param>
<param name="from">The reference frame that the position is in.</param>
<param name="to">The reference frame to covert the position to.</param>
<returns>The corresponding position, as a vector, in reference frame
<paramref name="to" />.</returns>
</doc>
```
#### Parameters 
```
position		
```
```
from		ReferenceFrame
```
```
to		ReferenceFrame
```
#### Return 
```
```
### TransformDirection

```
<doc>
<summary>
Converts a direction from one reference frame to another.
</summary>
<param name="direction">Direction, as a vector, in reference frame
<paramref name="from" />. </param>
<param name="from">The reference frame that the direction is in.</param>
<param name="to">The reference frame to covert the direction to.</param>
<returns>The corresponding direction, as a vector, in reference frame
<paramref name="to" />.</returns>
</doc>
```
#### Parameters 
```
direction		
```
```
from		ReferenceFrame
```
```
to		ReferenceFrame
```
#### Return 
```
```
### TransformRotation

```
<doc>
<summary>
Converts a rotation from one reference frame to another.
</summary>
<param name="rotation">Rotation, as a quaternion of the form <math>(x, y, z, w)</math>,
in reference frame <paramref name="from" />.</param>
<param name="from">The reference frame that the rotation is in.</param>
<param name="to">The reference frame to covert the rotation to.</param>
<returns>The corresponding rotation, as a quaternion of the form
<math>(x, y, z, w)</math>, in reference frame <paramref name="to" />.</returns>
</doc>
```
#### Parameters 
```
rotation		
```
```
from		ReferenceFrame
```
```
to		ReferenceFrame
```
#### Return 
```
```
### TransformVelocity

```
<doc>
<summary>
Converts a velocity (acting at the specified position) from one reference frame
to another. The position is required to take the relative angular velocity of the
reference frames into account.
</summary>
<param name="position">Position, as a vector, in reference frame
<paramref name="from" />.</param>
<param name="velocity">Velocity, as a vector that points in the direction of travel and
whose magnitude is the speed in meters per second, in reference frame
<paramref name="from" />.</param>
<param name="from">The reference frame that the position and velocity are in.</param>
<param name="to">The reference frame to covert the velocity to.</param>
<returns>The corresponding velocity, as a vector, in reference frame
<paramref name="to" />.</returns>
</doc>
```
#### Parameters 
```
position		
```
```
velocity		
```
```
from		ReferenceFrame
```
```
to		ReferenceFrame
```
#### Return 
```
```
### RaycastDistance

```
<doc>
<summary>
Cast a ray from a given position in a given direction, and return the distance to the hit point.
If no hit occurs, returns infinity.
</summary>
<param name="position">Position, as a vector, of the origin of the ray.</param>
<param name="direction">Direction of the ray, as a unit vector.</param>
<param name="referenceFrame">The reference frame that the position and direction are in.</param>
<returns>The distance to the hit, in meters, or infinity if there was no hit.</returns>
</doc>
```
#### Parameters 
```
position		
```
```
direction		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### RaycastPart

```
<doc>
<summary>
Cast a ray from a given position in a given direction, and return the part that it hits.
If no hit occurs, returns <c>null</c>.
</summary>
<param name="position">Position, as a vector, of the origin of the ray.</param>
<param name="direction">Direction of the ray, as a unit vector.</param>
<param name="referenceFrame">The reference frame that the position and direction are in.</param>
<returns>The part that was hit or <c>null</c> if there was no hit.</returns>
</doc>
```
#### Parameters 
```
position		
```
```
direction		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
Part```
### get_GameMode

```
<doc>
<summary>
The current mode the game is in.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
GameMode```
### get_Science

```
<doc>
<summary>
The current amount of science.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_Funds

```
<doc>
<summary>
The current amount of funds.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_Reputation

```
<doc>
<summary>
The current amount of reputation.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_ActiveVessel

```
<doc>
<summary>
The currently active vessel.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Vessel```
### set_ActiveVessel

```
<doc>
<summary>
The currently active vessel.
</summary>
</doc>
```
#### Parameters 
```
value		Vessel
```
#### Return 
```
```
### get_Vessels

```
<doc>
<summary>
A list of all the vessels in the game.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_Bodies

```
<doc>
<summary>
A dictionary of all celestial bodies (planets, moons, etc.) in the game,
keyed by the name of the body.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_TargetBody

```
<doc>
<summary>
The currently targeted celestial body.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
CelestialBody```
### set_TargetBody

```
<doc>
<summary>
The currently targeted celestial body.
</summary>
</doc>
```
#### Parameters 
```
value		CelestialBody
```
#### Return 
```
```
### get_TargetVessel

```
<doc>
<summary>
The currently targeted vessel.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Vessel```
### set_TargetVessel

```
<doc>
<summary>
The currently targeted vessel.
</summary>
</doc>
```
#### Parameters 
```
value		Vessel
```
#### Return 
```
```
### get_TargetDockingPort

```
<doc>
<summary>
The currently targeted docking port.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
DockingPort```
### set_TargetDockingPort

```
<doc>
<summary>
The currently targeted docking port.
</summary>
</doc>
```
#### Parameters 
```
value		DockingPort
```
#### Return 
```
```
### get_WaypointManager

```
<doc>
<summary>
The waypoint manager.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
WaypointManager```
### get_ContractManager

```
<doc>
<summary>
The contract manager.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
ContractManager```
### get_Camera

```
<doc>
<summary>
An object that can be used to control the camera.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Camera```
### get_UIVisible

```
<doc>
<summary>
Whether the UI is visible.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### set_UIVisible

```
<doc>
<summary>
Whether the UI is visible.
</summary>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
```
### get_Navball

```
<doc>
<summary>
Whether the navball is visible.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### set_Navball

```
<doc>
<summary>
Whether the navball is visible.
</summary>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
```
### get_UT

```
<doc>
<summary>
The current universal time in seconds.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_G

```
<doc>
<summary>
The value of the <a href="https://en.wikipedia.org/wiki/Gravitational_constant">
gravitational constant</a> G in <math>N(m/kg)^2</math>.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_WarpMode

```
<doc>
<summary>
The current time warp mode. Returns <see cref="M:SpaceCenter.WarpMode.None" /> if time
warp is not active, <see cref="M:SpaceCenter.WarpMode.Rails" /> if regular "on-rails" time warp
is active, or <see cref="M:SpaceCenter.WarpMode.Physics" /> if physical time warp is active.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
WarpMode```
### get_WarpRate

```
<doc>
<summary>
The current warp rate. This is the rate at which time is passing for
either on-rails or physical time warp. For example, a value of 10 means
time is passing 10x faster than normal. Returns 1 if time warp is not
active.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_WarpFactor

```
<doc>
<summary>
The current warp factor. This is the index of the rate at which time
is passing for either regular "on-rails" or physical time warp. Returns 0
if time warp is not active. When in on-rails time warp, this is equal to
<see cref="M:SpaceCenter.RailsWarpFactor" />, and in physics time warp, this is equal to
<see cref="M:SpaceCenter.PhysicsWarpFactor" />.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_RailsWarpFactor

```
<doc>
<summary>
The time warp rate, using regular "on-rails" time warp. A value between
0 and 7 inclusive. 0 means no time warp. Returns 0 if physical time warp
is active.

If requested time warp factor cannot be set, it will be set to the next
lowest possible value. For example, if the vessel is too close to a
planet. See <a href="https://wiki.kerbalspaceprogram.com/wiki/Time_warp">
the KSP wiki</a> for details.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### set_RailsWarpFactor

```
<doc>
<summary>
The time warp rate, using regular "on-rails" time warp. A value between
0 and 7 inclusive. 0 means no time warp. Returns 0 if physical time warp
is active.

If requested time warp factor cannot be set, it will be set to the next
lowest possible value. For example, if the vessel is too close to a
planet. See <a href="https://wiki.kerbalspaceprogram.com/wiki/Time_warp">
the KSP wiki</a> for details.
</summary>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
```
### get_PhysicsWarpFactor

```
<doc>
<summary>
The physical time warp rate. A value between 0 and 3 inclusive. 0 means
no time warp. Returns 0 if regular "on-rails" time warp is active.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### set_PhysicsWarpFactor

```
<doc>
<summary>
The physical time warp rate. A value between 0 and 3 inclusive. 0 means
no time warp. Returns 0 if regular "on-rails" time warp is active.
</summary>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
```
### get_MaximumRailsWarpFactor

```
<doc>
<summary>
The current maximum regular "on-rails" warp factor that can be set.
A value between 0 and 7 inclusive. See
<a href="https://wiki.kerbalspaceprogram.com/wiki/Time_warp">the KSP wiki</a>
for details.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_FARAvailable

```
<doc>
<summary>
Whether <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a> is installed.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### AutoPilot_Engage

```
<doc>
<summary>
Engage the auto-pilot.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_Disengage

```
<doc>
<summary>
Disengage the auto-pilot.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_Wait

```
<doc>
<summary>
Blocks until the vessel is pointing in the target direction and has
the target roll (if set). Throws an exception if the auto-pilot has not been engaged.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_TargetPitchAndHeading

```
<doc>
<summary>
Set target pitch and heading angles.
</summary>
<param name="pitch">Target pitch angle, in degrees between -90° and +90°.</param>
<param name="heading">Target heading angle, in degrees between 0° and 360°.</param>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
pitch		
```
```
heading		
```
#### Return 
```
```
### AutoPilot_get_Error

```
<doc>
<summary>
The error, in degrees, between the direction the ship has been asked
to point in and the direction it is pointing in. Throws an exception if the auto-pilot
has not been engaged and SAS is not enabled or is in stability assist mode.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_get_PitchError

```
<doc>
<summary>
The error, in degrees, between the vessels current and target pitch.
Throws an exception if the auto-pilot has not been engaged.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_get_HeadingError

```
<doc>
<summary>
The error, in degrees, between the vessels current and target heading.
Throws an exception if the auto-pilot has not been engaged.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_get_RollError

```
<doc>
<summary>
The error, in degrees, between the vessels current and target roll.
Throws an exception if the auto-pilot has not been engaged or no target roll is set.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_get_ReferenceFrame

```
<doc>
<summary>
The reference frame for the target direction (<see cref="M:SpaceCenter.AutoPilot.TargetDirection" />).
</summary>
<remarks>
An error will be thrown if this property is set to a reference frame that rotates with
the vessel being controlled, as it is impossible to rotate the vessel in such a
reference frame.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
ReferenceFrame```
### AutoPilot_set_ReferenceFrame

```
<doc>
<summary>
The reference frame for the target direction (<see cref="M:SpaceCenter.AutoPilot.TargetDirection" />).
</summary>
<remarks>
An error will be thrown if this property is set to a reference frame that rotates with
the vessel being controlled, as it is impossible to rotate the vessel in such a
reference frame.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		ReferenceFrame
```
#### Return 
```
```
### AutoPilot_get_TargetPitch

```
<doc>
<summary>
The target pitch, in degrees, between -90° and +90°.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_TargetPitch

```
<doc>
<summary>
The target pitch, in degrees, between -90° and +90°.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_TargetHeading

```
<doc>
<summary>
The target heading, in degrees, between 0° and 360°.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_TargetHeading

```
<doc>
<summary>
The target heading, in degrees, between 0° and 360°.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_TargetRoll

```
<doc>
<summary>
The target roll, in degrees. <c>NaN</c> if no target roll is set.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_TargetRoll

```
<doc>
<summary>
The target roll, in degrees. <c>NaN</c> if no target roll is set.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_TargetDirection

```
<doc>
<summary>
Direction vector corresponding to the target pitch and heading.
This is in the reference frame specified by <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_TargetDirection

```
<doc>
<summary>
Direction vector corresponding to the target pitch and heading.
This is in the reference frame specified by <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_SAS

```
<doc>
<summary>
The state of SAS.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.Control.SAS" /></remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_SAS

```
<doc>
<summary>
The state of SAS.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.Control.SAS" /></remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_SASMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SASMode" />.
These modes are equivalent to the mode buttons to the left of the navball that appear
when SAS is enabled.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.Control.SASMode" /></remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
SASMode```
### AutoPilot_set_SASMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SASMode" />.
These modes are equivalent to the mode buttons to the left of the navball that appear
when SAS is enabled.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.Control.SASMode" /></remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		SASMode
```
#### Return 
```
```
### AutoPilot_get_RollThreshold

```
<doc>
<summary>
The threshold at which the autopilot will try to match the target roll angle, if any.
Defaults to 5 degrees.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_RollThreshold

```
<doc>
<summary>
The threshold at which the autopilot will try to match the target roll angle, if any.
Defaults to 5 degrees.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_StoppingTime

```
<doc>
<summary>
The maximum amount of time that the vessel should need to come to a complete stop.
This determines the maximum angular velocity of the vessel.
A vector of three stopping times, in seconds, one for each of the pitch, roll
and yaw axes. Defaults to 0.5 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_StoppingTime

```
<doc>
<summary>
The maximum amount of time that the vessel should need to come to a complete stop.
This determines the maximum angular velocity of the vessel.
A vector of three stopping times, in seconds, one for each of the pitch, roll
and yaw axes. Defaults to 0.5 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_DecelerationTime

```
<doc>
<summary>
The time the vessel should take to come to a stop pointing in the target direction.
This determines the angular acceleration used to decelerate the vessel.
A vector of three times, in seconds, one for each of the pitch, roll and yaw axes.
Defaults to 5 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_DecelerationTime

```
<doc>
<summary>
The time the vessel should take to come to a stop pointing in the target direction.
This determines the angular acceleration used to decelerate the vessel.
A vector of three times, in seconds, one for each of the pitch, roll and yaw axes.
Defaults to 5 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_AttenuationAngle

```
<doc>
<summary>
The angle at which the autopilot considers the vessel to be pointing
close to the target.
This determines the midpoint of the target velocity attenuation function.
A vector of three angles, in degrees, one for each of the pitch, roll and yaw axes.
Defaults to 1° for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_AttenuationAngle

```
<doc>
<summary>
The angle at which the autopilot considers the vessel to be pointing
close to the target.
This determines the midpoint of the target velocity attenuation function.
A vector of three angles, in degrees, one for each of the pitch, roll and yaw axes.
Defaults to 1° for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_AutoTune

```
<doc>
<summary>
Whether the rotation rate controllers PID parameters should be automatically tuned
using the vessels moment of inertia and available torque. Defaults to <c>true</c>.
See <see cref="M:SpaceCenter.AutoPilot.TimeToPeak" /> and <see cref="M:SpaceCenter.AutoPilot.Overshoot" />.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_AutoTune

```
<doc>
<summary>
Whether the rotation rate controllers PID parameters should be automatically tuned
using the vessels moment of inertia and available torque. Defaults to <c>true</c>.
See <see cref="M:SpaceCenter.AutoPilot.TimeToPeak" /> and <see cref="M:SpaceCenter.AutoPilot.Overshoot" />.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_TimeToPeak

```
<doc>
<summary>
The target time to peak used to autotune the PID controllers.
A vector of three times, in seconds, for each of the pitch, roll and yaw axes.
Defaults to 3 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_TimeToPeak

```
<doc>
<summary>
The target time to peak used to autotune the PID controllers.
A vector of three times, in seconds, for each of the pitch, roll and yaw axes.
Defaults to 3 seconds for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_Overshoot

```
<doc>
<summary>
The target overshoot percentage used to autotune the PID controllers.
A vector of three values, between 0 and 1, for each of the pitch, roll and yaw axes.
Defaults to 0.01 for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_Overshoot

```
<doc>
<summary>
The target overshoot percentage used to autotune the PID controllers.
A vector of three values, between 0 and 1, for each of the pitch, roll and yaw axes.
Defaults to 0.01 for each axis.
</summary>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_PitchPIDGains

```
<doc>
<summary>
Gains for the pitch PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_PitchPIDGains

```
<doc>
<summary>
Gains for the pitch PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_RollPIDGains

```
<doc>
<summary>
Gains for the roll PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_RollPIDGains

```
<doc>
<summary>
Gains for the roll PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### AutoPilot_get_YawPIDGains

```
<doc>
<summary>
Gains for the yaw PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
#### Return 
```
```
### AutoPilot_set_YawPIDGains

```
<doc>
<summary>
Gains for the yaw PID controller.
</summary>
<remarks>
When <see cref="M:SpaceCenter.AutoPilot.AutoTune" /> is true, these values are updated automatically,
which will overwrite any manual changes.
</remarks>
</doc>
```
#### Parameters 
```
this		AutoPilot
```
```
value		
```
#### Return 
```
```
### Camera_get_Mode

```
<doc>
<summary>
The current mode of the camera.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
CameraMode```
### Camera_set_Mode

```
<doc>
<summary>
The current mode of the camera.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		CameraMode
```
#### Return 
```
```
### Camera_get_Pitch

```
<doc>
<summary>
The pitch of the camera, in degrees.
A value between <see cref="M:SpaceCenter.Camera.MinPitch" /> and <see cref="M:SpaceCenter.Camera.MaxPitch" /></summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_set_Pitch

```
<doc>
<summary>
The pitch of the camera, in degrees.
A value between <see cref="M:SpaceCenter.Camera.MinPitch" /> and <see cref="M:SpaceCenter.Camera.MaxPitch" /></summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		
```
#### Return 
```
```
### Camera_get_Heading

```
<doc>
<summary>
The heading of the camera, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_set_Heading

```
<doc>
<summary>
The heading of the camera, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		
```
#### Return 
```
```
### Camera_get_Distance

```
<doc>
<summary>
The distance from the camera to the subject, in meters.
A value between <see cref="M:SpaceCenter.Camera.MinDistance" /> and <see cref="M:SpaceCenter.Camera.MaxDistance" />.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_set_Distance

```
<doc>
<summary>
The distance from the camera to the subject, in meters.
A value between <see cref="M:SpaceCenter.Camera.MinDistance" /> and <see cref="M:SpaceCenter.Camera.MaxDistance" />.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		
```
#### Return 
```
```
### Camera_get_MinPitch

```
<doc>
<summary>
The minimum pitch of the camera.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_get_MaxPitch

```
<doc>
<summary>
The maximum pitch of the camera.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_get_MinDistance

```
<doc>
<summary>
Minimum distance from the camera to the subject, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_get_MaxDistance

```
<doc>
<summary>
Maximum distance from the camera to the subject, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_get_DefaultDistance

```
<doc>
<summary>
Default distance from the camera to the subject, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
```
### Camera_get_FocussedBody

```
<doc>
<summary>
In map mode, the celestial body that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a celestial body.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
CelestialBody```
### Camera_set_FocussedBody

```
<doc>
<summary>
In map mode, the celestial body that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a celestial body.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		CelestialBody
```
#### Return 
```
```
### Camera_get_FocussedVessel

```
<doc>
<summary>
In map mode, the vessel that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a vessel.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
Vessel```
### Camera_set_FocussedVessel

```
<doc>
<summary>
In map mode, the vessel that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a vessel.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		Vessel
```
#### Return 
```
```
### Camera_get_FocussedNode

```
<doc>
<summary>
In map mode, the maneuver node that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a maneuver node.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
#### Return 
```
Node```
### Camera_set_FocussedNode

```
<doc>
<summary>
In map mode, the maneuver node that the camera is focussed on.
Returns <c>null</c> if the camera is not focussed on a maneuver node.
Returns an error is the camera is not in map mode.
</summary>
</doc>
```
#### Parameters 
```
this		Camera
```
```
value		Node
```
#### Return 
```
```
### CelestialBody_SurfaceHeight

```
<doc>
<summary>
The height of the surface relative to mean sea level, in meters,
at the given position. When over water this is equal to 0.
</summary>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
#### Return 
```
```
### CelestialBody_BedrockHeight

```
<doc>
<summary>
The height of the surface relative to mean sea level, in meters,
at the given position. When over water, this is the height
of the sea-bed and is therefore  negative value.
</summary>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
#### Return 
```
```
### CelestialBody_MSLPosition

```
<doc>
<summary>
The position at mean sea level at the given latitude and longitude,
in the given reference frame.
</summary>
<returns>Position as a vector.</returns>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
<param name="referenceFrame">Reference frame for the returned position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_SurfacePosition

```
<doc>
<summary>
The position of the surface at the given latitude and longitude, in the given
reference frame. When over water, this is the position of the surface of the water.
</summary>
<returns>Position as a vector.</returns>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
<param name="referenceFrame">Reference frame for the returned position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_BedrockPosition

```
<doc>
<summary>
The position of the surface at the given latitude and longitude, in the given
reference frame. When over water, this is the position at the bottom of the sea-bed.
</summary>
<returns>Position as a vector.</returns>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
<param name="referenceFrame">Reference frame for the returned position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_PositionAtAltitude

```
<doc>
<summary>
The position at the given latitude, longitude and altitude, in the given reference frame.
</summary>
<returns>Position as a vector.</returns>
<param name="latitude">Latitude in degrees.</param>
<param name="longitude">Longitude in degrees.</param>
<param name="altitude">Altitude in meters above sea level.</param>
<param name="referenceFrame">Reference frame for the returned position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
```
altitude		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_LatitudeAtPosition

```
<doc>
<summary>
The latitude of the given position, in the given reference frame.
</summary>
<param name="position">Position as a vector.</param>
<param name="referenceFrame">Reference frame for the position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_LongitudeAtPosition

```
<doc>
<summary>
The longitude of the given position, in the given reference frame.
</summary>
<param name="position">Position as a vector.</param>
<param name="referenceFrame">Reference frame for the position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_AltitudeAtPosition

```
<doc>
<summary>
The altitude, in meters, of the given position in the given reference frame.
</summary>
<param name="position">Position as a vector.</param>
<param name="referenceFrame">Reference frame for the position vector.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_AtmosphericDensityAtPosition

```
<doc>
<summary>
The atmospheric density at the given position, in <math>kg/m^3</math>,
in the given reference frame.
</summary>
<param name="position">The position vector at which to measure the density.</param>
<param name="referenceFrame">Reference frame that the position vector is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_TemperatureAt

```
<doc>
<summary>
The temperature on the body at the given position, in the given reference frame.
</summary>
<param name="position">Position as a vector.</param>
<param name="referenceFrame">The reference frame that the position is in.</param>
<remarks>
This calculation is performed using the bodies current position, which means that
the value could be wrong if you want to know the temperature in the far future.
</remarks>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_DensityAt

```
<doc>
<summary>
Gets the air density, in <math>kg/m^3</math>, for the specified
altitude above sea level, in meters.
</summary>
<remarks>
This is an approximation, because actual calculations, taking sun exposure into account
to compute air temperature, require us to know the exact point on the body where the
density is to be computed (knowing the altitude is not enough).
However, the difference is small for high altitudes, so it makes very little difference
for trajectory prediction.
</remarks>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
altitude		
```
#### Return 
```
```
### CelestialBody_PressureAt

```
<doc>
<summary>
Gets the air pressure, in Pascals, for the specified
altitude above sea level, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
altitude		
```
#### Return 
```
```
### CelestialBody_BiomeAt

```
<doc>
<summary>
The biome at the given latitude and longitude, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
latitude		
```
```
longitude		
```
#### Return 
```
```
### CelestialBody_Position

```
<doc>
<summary>
The position of the center of the body, in the specified reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_Velocity

```
<doc>
<summary>
The linear velocity of the body, in the specified reference frame.
</summary>
<returns>The velocity as a vector. The vector points in the direction of travel,
and its magnitude is the speed of the body in meters per second.</returns>
<param name="referenceFrame">The reference frame that the returned
velocity vector is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_Rotation

```
<doc>
<summary>
The rotation of the body, in the specified reference frame.
</summary>
<returns>The rotation as a quaternion of the form <math>(x, y, z, w)</math>.</returns>
<param name="referenceFrame">The reference frame that the returned
rotation is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_Direction

```
<doc>
<summary>
The direction in which the north pole of the celestial body is pointing,
in the specified reference frame.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_AngularVelocity

```
<doc>
<summary>
The angular velocity of the body in the specified reference frame.
</summary>
<returns>The angular velocity as a vector. The magnitude of the vector is the rotational
speed of the body, in radians per second. The direction of the vector indicates the axis
of rotation, using the right-hand rule.</returns>
<param name="referenceFrame">The reference frame the returned
angular velocity is in.</param>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### CelestialBody_get_Name

```
<doc>
<summary>
The name of the body.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_Satellites

```
<doc>
<summary>
A list of celestial bodies that are in orbit around this celestial body.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_Mass

```
<doc>
<summary>
The mass of the body, in kilograms.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_GravitationalParameter

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Standard_gravitational_parameter">standard
gravitational parameter</a> of the body in <math>m^3s^{-2}</math>.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_SurfaceGravity

```
<doc>
<summary>
The acceleration due to gravity at sea level (mean altitude) on the body,
in <math>m/s^2</math>.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_RotationalPeriod

```
<doc>
<summary>
The sidereal rotational period of the body, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_RotationalSpeed

```
<doc>
<summary>
The rotational speed of the body, in radians per second.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_RotationAngle

```
<doc>
<summary>
The current rotation angle of the body, in radians.
A value between 0 and <math>2\pi</math></summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_InitialRotation

```
<doc>
<summary>
The initial rotation angle of the body (at UT 0), in radians.
A value between 0 and <math>2\pi</math></summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_EquatorialRadius

```
<doc>
<summary>
The equatorial radius of the body, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_SphereOfInfluence

```
<doc>
<summary>
The radius of the sphere of influence of the body, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_Orbit

```
<doc>
<summary>
The orbit of the body.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
Orbit```
### CelestialBody_get_HasAtmosphere

```
<doc>
<summary><c>true</c> if the body has an atmosphere.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_AtmosphereDepth

```
<doc>
<summary>
The depth of the atmosphere, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_HasAtmosphericOxygen

```
<doc>
<summary><c>true</c> if there is oxygen in the atmosphere, required for air-breathing engines.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_Biomes

```
<doc>
<summary>
The biomes present on this body.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_FlyingHighAltitudeThreshold

```
<doc>
<summary>
The altitude, in meters, above which a vessel is considered to be
flying "high" when doing science.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_SpaceHighAltitudeThreshold

```
<doc>
<summary>
The altitude, in meters, above which a vessel is considered to be
in "high" space when doing science.
</summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
```
### CelestialBody_get_ReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the celestial body.
<list type="bullet"><item><description>The origin is at the center of the body.
</description></item><item><description>The axes rotate with the body.</description></item><item><description>The x-axis points from the center of the body
towards the intersection of the prime meridian and equator (the
position at 0° longitude, 0° latitude).</description></item><item><description>The y-axis points from the center of the body
towards the north pole.</description></item><item><description>The z-axis points from the center of the body
towards the equator at 90°E longitude.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
ReferenceFrame```
### CelestialBody_get_NonRotatingReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to this celestial body, and
orientated in a fixed direction (it does not rotate with the body).
<list type="bullet"><item><description>The origin is at the center of the body.</description></item><item><description>The axes do not rotate.</description></item><item><description>The x-axis points in an arbitrary direction through the
equator.</description></item><item><description>The y-axis points from the center of the body towards
the north pole.</description></item><item><description>The z-axis points in an arbitrary direction through the
equator.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
ReferenceFrame```
### CelestialBody_get_OrbitalReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to this celestial body, but
orientated with the body's orbital prograde/normal/radial directions.
<list type="bullet"><item><description>The origin is at the center of the body.
</description></item><item><description>The axes rotate with the orbital prograde/normal/radial
directions.</description></item><item><description>The x-axis points in the orbital anti-radial direction.
</description></item><item><description>The y-axis points in the orbital prograde direction.
</description></item><item><description>The z-axis points in the orbital normal direction.
</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		CelestialBody
```
#### Return 
```
ReferenceFrame```
### CommLink_get_Type

```
<doc>
<summary>
The type of link.
</summary>
</doc>
```
#### Parameters 
```
this		CommLink
```
#### Return 
```
CommLinkType```
### CommLink_get_SignalStrength

```
<doc>
<summary>
Signal strength of the link.
</summary>
</doc>
```
#### Parameters 
```
this		CommLink
```
#### Return 
```
```
### CommLink_get_Start

```
<doc>
<summary>
Start point of the link.
</summary>
</doc>
```
#### Parameters 
```
this		CommLink
```
#### Return 
```
CommNode```
### CommLink_get_End

```
<doc>
<summary>
Start point of the link.
</summary>
</doc>
```
#### Parameters 
```
this		CommLink
```
#### Return 
```
CommNode```
### CommNode_get_Name

```
<doc>
<summary>
Name of the communication node.
</summary>
</doc>
```
#### Parameters 
```
this		CommNode
```
#### Return 
```
```
### CommNode_get_IsHome

```
<doc>
<summary>
Whether the communication node is on Kerbin.
</summary>
</doc>
```
#### Parameters 
```
this		CommNode
```
#### Return 
```
```
### CommNode_get_IsControlPoint

```
<doc>
<summary>
Whether the communication node is a control point, for example a manned vessel.
</summary>
</doc>
```
#### Parameters 
```
this		CommNode
```
#### Return 
```
```
### CommNode_get_IsVessel

```
<doc>
<summary>
Whether the communication node is a vessel.
</summary>
</doc>
```
#### Parameters 
```
this		CommNode
```
#### Return 
```
```
### CommNode_get_Vessel

```
<doc>
<summary>
The vessel for this communication node.
</summary>
</doc>
```
#### Parameters 
```
this		CommNode
```
#### Return 
```
Vessel```
### Comms_get_CanCommunicate

```
<doc>
<summary>
Whether the vessel can communicate with KSC.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_CanTransmitScience

```
<doc>
<summary>
Whether the vessel can transmit science data to KSC.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_SignalStrength

```
<doc>
<summary>
Signal strength to KSC.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_SignalDelay

```
<doc>
<summary>
Signal delay to KSC in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_Power

```
<doc>
<summary>
The combined power of all active antennae on the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Comms_get_ControlPath

```
<doc>
<summary>
The communication path used to control the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Comms
```
#### Return 
```
```
### Contract_Cancel

```
<doc>
<summary>
Cancel an active contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_Accept

```
<doc>
<summary>
Accept an offered contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_Decline

```
<doc>
<summary>
Decline an offered contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Type

```
<doc>
<summary>
Type of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Title

```
<doc>
<summary>
Title of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Description

```
<doc>
<summary>
Description of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Notes

```
<doc>
<summary>
Notes for the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Synopsis

```
<doc>
<summary>
Synopsis for the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Keywords

```
<doc>
<summary>
Keywords for the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_State

```
<doc>
<summary>
State of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
ContractState```
### Contract_get_Active

```
<doc>
<summary>
Whether the contract is active.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Failed

```
<doc>
<summary>
Whether the contract has been failed.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Seen

```
<doc>
<summary>
Whether the contract has been seen.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Read

```
<doc>
<summary>
Whether the contract has been read.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_CanBeCanceled

```
<doc>
<summary>
Whether the contract can be canceled.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_CanBeDeclined

```
<doc>
<summary>
Whether the contract can be declined.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_CanBeFailed

```
<doc>
<summary>
Whether the contract can be failed.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_FundsAdvance

```
<doc>
<summary>
Funds received when accepting the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_FundsCompletion

```
<doc>
<summary>
Funds received on completion of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_FundsFailure

```
<doc>
<summary>
Funds lost if the contract is failed.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_ReputationCompletion

```
<doc>
<summary>
Reputation gained on completion of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_ReputationFailure

```
<doc>
<summary>
Reputation lost if the contract is failed.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_ScienceCompletion

```
<doc>
<summary>
Science gained on completion of the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### Contract_get_Parameters

```
<doc>
<summary>
Parameters for the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Contract
```
#### Return 
```
```
### ContractManager_get_Types

```
<doc>
<summary>
A list of all contract types.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractManager_get_AllContracts

```
<doc>
<summary>
A list of all contracts.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractManager_get_ActiveContracts

```
<doc>
<summary>
A list of all active contracts.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractManager_get_OfferedContracts

```
<doc>
<summary>
A list of all offered, but unaccepted, contracts.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractManager_get_CompletedContracts

```
<doc>
<summary>
A list of all completed contracts.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractManager_get_FailedContracts

```
<doc>
<summary>
A list of all failed contracts.
</summary>
</doc>
```
#### Parameters 
```
this		ContractManager
```
#### Return 
```
```
### ContractParameter_get_Title

```
<doc>
<summary>
Title of the parameter.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_Notes

```
<doc>
<summary>
Notes for the parameter.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_Children

```
<doc>
<summary>
Child contract parameters.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_Completed

```
<doc>
<summary>
Whether the parameter has been completed.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_Failed

```
<doc>
<summary>
Whether the parameter has been failed.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_Optional

```
<doc>
<summary>
Whether the contract parameter is optional.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_FundsCompletion

```
<doc>
<summary>
Funds received on completion of the contract parameter.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_FundsFailure

```
<doc>
<summary>
Funds lost if the contract parameter is failed.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_ReputationCompletion

```
<doc>
<summary>
Reputation gained on completion of the contract parameter.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_ReputationFailure

```
<doc>
<summary>
Reputation lost if the contract parameter is failed.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### ContractParameter_get_ScienceCompletion

```
<doc>
<summary>
Science gained on completion of the contract parameter.
</summary>
</doc>
```
#### Parameters 
```
this		ContractParameter
```
#### Return 
```
```
### Control_ActivateNextStage

```
<doc>
<summary>
Activates the next stage. Equivalent to pressing the space bar in-game.
</summary>
<returns>A list of vessel objects that are jettisoned from the active vessel.</returns>
<remarks>
When called, the active vessel may change. It is therefore possible that,
after calling this function, the object(s) returned by previous call(s) to
<see cref="M:SpaceCenter.ActiveVessel" /> no longer refer to the active vessel.
</remarks>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_GetActionGroup

```
<doc>
<summary>
Returns <c>true</c> if the given action group is enabled.
</summary>
<param name="group">
A number between 0 and 9 inclusive,
or between 0 and 250 inclusive when the <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/67235-122dec1016-action-groups-extended-250-action-groups-in-flight-editing-now-kosremotetech/">Extended Action Groups mod</a> is installed.
</param>
</doc>
```
#### Parameters 
```
this		Control
```
```
group		
```
#### Return 
```
```
### Control_SetActionGroup

```
<doc>
<summary>
Sets the state of the given action group.
</summary>
<param name="group">
A number between 0 and 9 inclusive,
or between 0 and 250 inclusive when the <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/67235-122dec1016-action-groups-extended-250-action-groups-in-flight-editing-now-kosremotetech/">Extended Action Groups mod</a> is installed.
</param>
<param name="state"></param>
</doc>
```
#### Parameters 
```
this		Control
```
```
group		
```
```
state		
```
#### Return 
```
```
### Control_ToggleActionGroup

```
<doc>
<summary>
Toggles the state of the given action group.
</summary>
<param name="group">
A number between 0 and 9 inclusive,
or between 0 and 250 inclusive when the <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/67235-122dec1016-action-groups-extended-250-action-groups-in-flight-editing-now-kosremotetech/">Extended Action Groups mod</a> is installed.
</param>
</doc>
```
#### Parameters 
```
this		Control
```
```
group		
```
#### Return 
```
```
### Control_AddNode

```
<doc>
<summary>
Creates a maneuver node at the given universal time, and returns a
<see cref="T:SpaceCenter.Node" /> object that can be used to modify it.
Optionally sets the magnitude of the delta-v for the maneuver node
in the prograde, normal and radial directions.
</summary>
<param name="ut">Universal time of the maneuver node.</param>
<param name="prograde">Delta-v in the prograde direction.</param>
<param name="normal">Delta-v in the normal direction.</param>
<param name="radial">Delta-v in the radial direction.</param>
</doc>
```
#### Parameters 
```
this		Control
```
```
ut		
```
```
prograde		
```
```
normal		
```
```
radial		
```
#### Return 
```
Node```
### Control_RemoveNodes

```
<doc>
<summary>
Remove all maneuver nodes.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_get_State

```
<doc>
<summary>
The control state of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
ControlState```
### Control_get_Source

```
<doc>
<summary>
The source of the vessels control, for example by a kerbal or a probe core.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
ControlSource```
### Control_get_SAS

```
<doc>
<summary>
The state of SAS.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.AutoPilot.SAS" /></remarks>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_SAS

```
<doc>
<summary>
The state of SAS.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.AutoPilot.SAS" /></remarks>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_SASMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SASMode" />.
These modes are equivalent to the mode buttons to
the left of the navball that appear when SAS is enabled.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.AutoPilot.SASMode" /></remarks>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
SASMode```
### Control_set_SASMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SASMode" />.
These modes are equivalent to the mode buttons to
the left of the navball that appear when SAS is enabled.
</summary>
<remarks>Equivalent to <see cref="M:SpaceCenter.AutoPilot.SASMode" /></remarks>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		SASMode
```
#### Return 
```
```
### Control_get_SpeedMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SpeedMode" /> of the navball.
This is the mode displayed next to the speed at the top of the navball.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
SpeedMode```
### Control_set_SpeedMode

```
<doc>
<summary>
The current <see cref="T:SpaceCenter.SpeedMode" /> of the navball.
This is the mode displayed next to the speed at the top of the navball.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		SpeedMode
```
#### Return 
```
```
### Control_get_RCS

```
<doc>
<summary>
The state of RCS.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_RCS

```
<doc>
<summary>
The state of RCS.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_ReactionWheels

```
<doc>
<summary>
Returns whether all reactive wheels on the vessel are active,
and sets the active state of all reaction wheels.
See <see cref="M:SpaceCenter.ReactionWheel.Active" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_ReactionWheels

```
<doc>
<summary>
Returns whether all reactive wheels on the vessel are active,
and sets the active state of all reaction wheels.
See <see cref="M:SpaceCenter.ReactionWheel.Active" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Gear

```
<doc>
<summary>
The state of the landing gear/legs.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Gear

```
<doc>
<summary>
The state of the landing gear/legs.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Legs

```
<doc>
<summary>
Returns whether all landing legs on the vessel are deployed,
and sets the deployment state of all landing legs.
Does not include wheels (for example landing gear).
See <see cref="M:SpaceCenter.Leg.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Legs

```
<doc>
<summary>
Returns whether all landing legs on the vessel are deployed,
and sets the deployment state of all landing legs.
Does not include wheels (for example landing gear).
See <see cref="M:SpaceCenter.Leg.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Wheels

```
<doc>
<summary>
Returns whether all wheels on the vessel are deployed,
and sets the deployment state of all wheels.
Does not include landing legs.
See <see cref="M:SpaceCenter.Wheel.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Wheels

```
<doc>
<summary>
Returns whether all wheels on the vessel are deployed,
and sets the deployment state of all wheels.
Does not include landing legs.
See <see cref="M:SpaceCenter.Wheel.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Lights

```
<doc>
<summary>
The state of the lights.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Lights

```
<doc>
<summary>
The state of the lights.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Brakes

```
<doc>
<summary>
The state of the wheel brakes.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Brakes

```
<doc>
<summary>
The state of the wheel brakes.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Antennas

```
<doc>
<summary>
Returns whether all antennas on the vessel are deployed,
and sets the deployment state of all antennas.
See <see cref="M:SpaceCenter.Antenna.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Antennas

```
<doc>
<summary>
Returns whether all antennas on the vessel are deployed,
and sets the deployment state of all antennas.
See <see cref="M:SpaceCenter.Antenna.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_CargoBays

```
<doc>
<summary>
Returns whether any of the cargo bays on the vessel are open,
and sets the open state of all cargo bays.
See <see cref="M:SpaceCenter.CargoBay.Open" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_CargoBays

```
<doc>
<summary>
Returns whether any of the cargo bays on the vessel are open,
and sets the open state of all cargo bays.
See <see cref="M:SpaceCenter.CargoBay.Open" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Intakes

```
<doc>
<summary>
Returns whether all of the air intakes on the vessel are open,
and sets the open state of all air intakes.
See <see cref="M:SpaceCenter.Intake.Open" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Intakes

```
<doc>
<summary>
Returns whether all of the air intakes on the vessel are open,
and sets the open state of all air intakes.
See <see cref="M:SpaceCenter.Intake.Open" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Parachutes

```
<doc>
<summary>
Returns whether all parachutes on the vessel are deployed,
and sets the deployment state of all parachutes.
Cannot be set to <c>false</c>.
See <see cref="M:SpaceCenter.Parachute.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Parachutes

```
<doc>
<summary>
Returns whether all parachutes on the vessel are deployed,
and sets the deployment state of all parachutes.
Cannot be set to <c>false</c>.
See <see cref="M:SpaceCenter.Parachute.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Radiators

```
<doc>
<summary>
Returns whether all radiators on the vessel are deployed,
and sets the deployment state of all radiators.
See <see cref="M:SpaceCenter.Radiator.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Radiators

```
<doc>
<summary>
Returns whether all radiators on the vessel are deployed,
and sets the deployment state of all radiators.
See <see cref="M:SpaceCenter.Radiator.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_ResourceHarvesters

```
<doc>
<summary>
Returns whether all of the resource harvesters on the vessel are deployed,
and sets the deployment state of all resource harvesters.
See <see cref="M:SpaceCenter.ResourceHarvester.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_ResourceHarvesters

```
<doc>
<summary>
Returns whether all of the resource harvesters on the vessel are deployed,
and sets the deployment state of all resource harvesters.
See <see cref="M:SpaceCenter.ResourceHarvester.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_ResourceHarvestersActive

```
<doc>
<summary>
Returns whether any of the resource harvesters on the vessel are active,
and sets the active state of all resource harvesters.
See <see cref="M:SpaceCenter.ResourceHarvester.Active" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_ResourceHarvestersActive

```
<doc>
<summary>
Returns whether any of the resource harvesters on the vessel are active,
and sets the active state of all resource harvesters.
See <see cref="M:SpaceCenter.ResourceHarvester.Active" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_SolarPanels

```
<doc>
<summary>
Returns whether all solar panels on the vessel are deployed,
and sets the deployment state of all solar panels.
See <see cref="M:SpaceCenter.SolarPanel.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_SolarPanels

```
<doc>
<summary>
Returns whether all solar panels on the vessel are deployed,
and sets the deployment state of all solar panels.
See <see cref="M:SpaceCenter.SolarPanel.Deployed" />.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Abort

```
<doc>
<summary>
The state of the abort action group.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Abort

```
<doc>
<summary>
The state of the abort action group.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Throttle

```
<doc>
<summary>
The state of the throttle. A value between 0 and 1.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Throttle

```
<doc>
<summary>
The state of the throttle. A value between 0 and 1.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_InputMode

```
<doc>
<summary>
Sets the behavior of the pitch, yaw, roll and translation control inputs.
When set to additive, these inputs are added to the vessels current inputs.
This mode is the default.
When set to override, these inputs (if non-zero) override the vessels inputs.
This mode prevents keyboard control, or SAS, from interfering with the controls when
they are set.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
ControlInputMode```
### Control_set_InputMode

```
<doc>
<summary>
Sets the behavior of the pitch, yaw, roll and translation control inputs.
When set to additive, these inputs are added to the vessels current inputs.
This mode is the default.
When set to override, these inputs (if non-zero) override the vessels inputs.
This mode prevents keyboard control, or SAS, from interfering with the controls when
they are set.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		ControlInputMode
```
#### Return 
```
```
### Control_get_Pitch

```
<doc>
<summary>
The state of the pitch control.
A value between -1 and 1.
Equivalent to the w and s keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Pitch

```
<doc>
<summary>
The state of the pitch control.
A value between -1 and 1.
Equivalent to the w and s keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Yaw

```
<doc>
<summary>
The state of the yaw control.
A value between -1 and 1.
Equivalent to the a and d keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Yaw

```
<doc>
<summary>
The state of the yaw control.
A value between -1 and 1.
Equivalent to the a and d keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Roll

```
<doc>
<summary>
The state of the roll control.
A value between -1 and 1.
Equivalent to the q and e keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Roll

```
<doc>
<summary>
The state of the roll control.
A value between -1 and 1.
Equivalent to the q and e keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Forward

```
<doc>
<summary>
The state of the forward translational control.
A value between -1 and 1.
Equivalent to the h and n keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Forward

```
<doc>
<summary>
The state of the forward translational control.
A value between -1 and 1.
Equivalent to the h and n keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Up

```
<doc>
<summary>
The state of the up translational control.
A value between -1 and 1.
Equivalent to the i and k keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Up

```
<doc>
<summary>
The state of the up translational control.
A value between -1 and 1.
Equivalent to the i and k keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_Right

```
<doc>
<summary>
The state of the right translational control.
A value between -1 and 1.
Equivalent to the j and l keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_Right

```
<doc>
<summary>
The state of the right translational control.
A value between -1 and 1.
Equivalent to the j and l keys.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_WheelThrottle

```
<doc>
<summary>
The state of the wheel throttle.
A value between -1 and 1.
A value of 1 rotates the wheels forwards, a value of -1 rotates
the wheels backwards.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_WheelThrottle

```
<doc>
<summary>
The state of the wheel throttle.
A value between -1 and 1.
A value of 1 rotates the wheels forwards, a value of -1 rotates
the wheels backwards.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_WheelSteering

```
<doc>
<summary>
The state of the wheel steering.
A value between -1 and 1.
A value of 1 steers to the left, and a value of -1 steers to the right.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_set_WheelSteering

```
<doc>
<summary>
The state of the wheel steering.
A value between -1 and 1.
A value of 1 steers to the left, and a value of -1 steers to the right.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
```
value		
```
#### Return 
```
```
### Control_get_CurrentStage

```
<doc>
<summary>
The current stage of the vessel. Corresponds to the stage number in
the in-game UI.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### Control_get_Nodes

```
<doc>
<summary>
Returns a list of all existing maneuver nodes, ordered by time from first to last.
</summary>
</doc>
```
#### Parameters 
```
this		Control
```
#### Return 
```
```
### CrewMember_get_Name

```
<doc>
<summary>
The crew members name.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Name

```
<doc>
<summary>
The crew members name.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### CrewMember_get_Type

```
<doc>
<summary>
The type of crew member.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
CrewMemberType```
### CrewMember_get_OnMission

```
<doc>
<summary>
Whether the crew member is on a mission.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_get_Courage

```
<doc>
<summary>
The crew members courage.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Courage

```
<doc>
<summary>
The crew members courage.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### CrewMember_get_Stupidity

```
<doc>
<summary>
The crew members stupidity.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Stupidity

```
<doc>
<summary>
The crew members stupidity.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### CrewMember_get_Experience

```
<doc>
<summary>
The crew members experience.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Experience

```
<doc>
<summary>
The crew members experience.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### CrewMember_get_Badass

```
<doc>
<summary>
Whether the crew member is a badass.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Badass

```
<doc>
<summary>
Whether the crew member is a badass.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### CrewMember_get_Veteran

```
<doc>
<summary>
Whether the crew member is a veteran.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
#### Return 
```
```
### CrewMember_set_Veteran

```
<doc>
<summary>
Whether the crew member is a veteran.
</summary>
</doc>
```
#### Parameters 
```
this		CrewMember
```
```
value		
```
#### Return 
```
```
### Flight_SimulateAerodynamicForceAt

```
<doc>
<summary>
Simulate and return the total aerodynamic forces acting on the vessel,
if it where to be traveling with the given velocity at the given position in the
atmosphere of the given celestial body.
</summary>
<returns>A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
```
body		CelestialBody
```
```
position		
```
```
velocity		
```
#### Return 
```
```
### Flight_get_GForce

```
<doc>
<summary>
The current G force acting on the vessel in <math>m/s^2</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_MeanAltitude

```
<doc>
<summary>
The altitude above sea level, in meters.
Measured from the center of mass of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_SurfaceAltitude

```
<doc>
<summary>
The altitude above the surface of the body or sea level, whichever is closer, in meters.
Measured from the center of mass of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_BedrockAltitude

```
<doc>
<summary>
The altitude above the surface of the body, in meters. When over water, this is the altitude above the sea floor.
Measured from the center of mass of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Elevation

```
<doc>
<summary>
The elevation of the terrain under the vessel, in meters. This is the height of the terrain above sea level,
and is negative when the vessel is over the sea.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Latitude

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Latitude">latitude</a> of the vessel for the body being orbited, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Longitude

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Longitude">longitude</a> of the vessel for the body being orbited, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Velocity

```
<doc>
<summary>
The velocity of the vessel, in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The velocity as a vector. The vector points in the direction of travel,
and its magnitude is the speed of the vessel in meters per second.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Speed

```
<doc>
<summary>
The speed of the vessel in meters per second,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_HorizontalSpeed

```
<doc>
<summary>
The horizontal speed of the vessel in meters per second,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_VerticalSpeed

```
<doc>
<summary>
The vertical speed of the vessel in meters per second,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_CenterOfMass

```
<doc>
<summary>
The position of the center of mass of the vessel,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" /></summary>
<returns>The position as a vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Rotation

```
<doc>
<summary>
The rotation of the vessel, in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" /></summary>
<returns>The rotation as a quaternion of the form <math>(x, y, z, w)</math>.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Direction

```
<doc>
<summary>
The direction that the vessel is pointing in,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Pitch

```
<doc>
<summary>
The pitch of the vessel relative to the horizon, in degrees.
A value between -90° and +90°.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Heading

```
<doc>
<summary>
The heading of the vessel (its angle relative to north), in degrees.
A value between 0° and 360°.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Roll

```
<doc>
<summary>
The roll of the vessel relative to the horizon, in degrees.
A value between -180° and +180°.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Prograde

```
<doc>
<summary>
The prograde direction of the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Retrograde

```
<doc>
<summary>
The retrograde direction of the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Normal

```
<doc>
<summary>
The direction normal to the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_AntiNormal

```
<doc>
<summary>
The direction opposite to the normal of the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Radial

```
<doc>
<summary>
The radial direction of the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_AntiRadial

```
<doc>
<summary>
The direction opposite to the radial direction of the vessels orbit,
in the reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The direction as a unit vector.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_AtmosphereDensity

```
<doc>
<summary>
The current density of the atmosphere around the vessel, in <math>kg/m^3</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_DynamicPressure

```
<doc>
<summary>
The dynamic pressure acting on the vessel, in Pascals. This is a measure of the
strength of the aerodynamic forces. It is equal to
<math>\frac{1}{2} . \mbox{air density} . \mbox{velocity}^2</math>.
It is commonly denoted <math>Q</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_StaticPressureAtMSL

```
<doc>
<summary>
The static atmospheric pressure at mean sea level, in Pascals.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_StaticPressure

```
<doc>
<summary>
The static atmospheric pressure acting on the vessel, in Pascals.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_AerodynamicForce

```
<doc>
<summary>
The total aerodynamic forces acting on the vessel,
in reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Lift

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Aerodynamic_force">aerodynamic lift</a>
currently acting on the vessel.
</summary>
<returns>A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Drag

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Aerodynamic_force">aerodynamic drag</a> currently acting on the vessel.
</summary>
<returns>A vector pointing in the direction of the force, with its magnitude
equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_SpeedOfSound

```
<doc>
<summary>
The speed of sound, in the atmosphere around the vessel, in <math>m/s</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_Mach

```
<doc>
<summary>
The speed of the vessel, in multiples of the speed of sound.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_ReynoldsNumber

```
<doc>
<summary>
The vessels Reynolds number.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_TrueAirSpeed

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/True_airspeed">true air speed</a>
of the vessel, in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_EquivalentAirSpeed

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Equivalent_airspeed">equivalent air speed</a>
of the vessel, in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_TerminalVelocity

```
<doc>
<summary>
An estimate of the current terminal velocity of the vessel, in meters per second.
This is the speed at which the drag forces cancel out the force of gravity.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_AngleOfAttack

```
<doc>
<summary>
The pitch angle between the orientation of the vessel and its velocity vector,
in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_SideslipAngle

```
<doc>
<summary>
The yaw angle between the orientation of the vessel and its velocity vector, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_TotalAirTemperature

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Total_air_temperature">total air temperature</a>
of the atmosphere around the vessel, in Kelvin.
This includes the <see cref="M:SpaceCenter.Flight.StaticAirTemperature" /> and the vessel's kinetic energy.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_StaticAirTemperature

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Total_air_temperature">static (ambient)
temperature</a> of the atmosphere around the vessel, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_StallFraction

```
<doc>
<summary>
The current amount of stall, between 0 and 1. A value greater than 0.005 indicates
a minor stall and a value greater than 0.5 indicates a large-scale stall.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_DragCoefficient

```
<doc>
<summary>
The coefficient of drag. This is the amount of drag produced by the vessel.
It depends on air speed, air density and wing area.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_LiftCoefficient

```
<doc>
<summary>
The coefficient of lift. This is the amount of lift produced by the vessel, and
depends on air speed, air density and wing area.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_BallisticCoefficient

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Ballistic_coefficient">ballistic coefficient</a>.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Flight_get_ThrustSpecificFuelConsumption

```
<doc>
<summary>
The thrust specific fuel consumption for the jet engines on the vessel. This is a
measure of the efficiency of the engines, with a lower value indicating a more
efficient vessel. This value is the number of Newtons of fuel that are burned,
per hour, to produce one newton of thrust.
</summary>
<remarks>
Requires <a href="https://forum.kerbalspaceprogram.com/index.php?/topic/19321-130-ferram-aerospace-research-v0159-liebe-82117/">Ferram Aerospace Research</a>.
</remarks>
</doc>
```
#### Parameters 
```
this		Flight
```
#### Return 
```
```
### Node_BurnVector

```
<doc>
<summary>
Returns the burn vector for the maneuver node.
</summary>
<param name="referenceFrame">The reference frame that the returned vector is in.
Defaults to <see cref="M:SpaceCenter.Vessel.OrbitalReferenceFrame" />.</param>
<returns>A vector whose direction is the direction of the maneuver node burn, and
magnitude is the delta-v of the burn in meters per second.
</returns>
<remarks>
Does not change when executing the maneuver node. See <see cref="M:SpaceCenter.Node.RemainingBurnVector" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Node
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Node_RemainingBurnVector

```
<doc>
<summary>
Returns the remaining burn vector for the maneuver node.
</summary>
<param name="referenceFrame">The reference frame that the returned vector is in.
Defaults to <see cref="M:SpaceCenter.Vessel.OrbitalReferenceFrame" />.</param>
<returns>A vector whose direction is the direction of the maneuver node burn, and
magnitude is the delta-v of the burn in meters per second.
</returns>
<remarks>
Changes as the maneuver node is executed. See <see cref="M:SpaceCenter.Node.BurnVector" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Node
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Node_Remove

```
<doc>
<summary>
Removes the maneuver node.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_Position

```
<doc>
<summary>
The position vector of the maneuver node in the given reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Node
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Node_Direction

```
<doc>
<summary>
The direction of the maneuver nodes burn.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		Node
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Node_get_Prograde

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the prograde direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_set_Prograde

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the prograde direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
```
value		
```
#### Return 
```
```
### Node_get_Normal

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the normal direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_set_Normal

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the normal direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
```
value		
```
#### Return 
```
```
### Node_get_Radial

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the radial direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_set_Radial

```
<doc>
<summary>
The magnitude of the maneuver nodes delta-v in the radial direction,
in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
```
value		
```
#### Return 
```
```
### Node_get_DeltaV

```
<doc>
<summary>
The delta-v of the maneuver node, in meters per second.
</summary>
<remarks>
Does not change when executing the maneuver node. See <see cref="M:SpaceCenter.Node.RemainingDeltaV" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_set_DeltaV

```
<doc>
<summary>
The delta-v of the maneuver node, in meters per second.
</summary>
<remarks>
Does not change when executing the maneuver node. See <see cref="M:SpaceCenter.Node.RemainingDeltaV" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Node
```
```
value		
```
#### Return 
```
```
### Node_get_RemainingDeltaV

```
<doc>
<summary>
Gets the remaining delta-v of the maneuver node, in meters per second. Changes as the
node is executed. This is equivalent to the delta-v reported in-game.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_get_UT

```
<doc>
<summary>
The universal time at which the maneuver will occur, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_set_UT

```
<doc>
<summary>
The universal time at which the maneuver will occur, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
```
value		
```
#### Return 
```
```
### Node_get_TimeTo

```
<doc>
<summary>
The time until the maneuver node will be encountered, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
```
### Node_get_Orbit

```
<doc>
<summary>
The orbit that results from executing the maneuver node.
</summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
Orbit```
### Node_get_ReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the maneuver node's burn.
<list type="bullet"><item><description>The origin is at the position of the maneuver node.</description></item><item><description>The y-axis points in the direction of the burn.</description></item><item><description>The x-axis and z-axis point in arbitrary but fixed directions.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
ReferenceFrame```
### Node_get_OrbitalReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the maneuver node, and
orientated with the orbital prograde/normal/radial directions of the
original orbit at the maneuver node's position.
<list type="bullet"><item><description>The origin is at the position of the maneuver node.</description></item><item><description>The x-axis points in the orbital anti-radial direction of the original
orbit, at the position of the maneuver node.</description></item><item><description>The y-axis points in the orbital prograde direction of the original
orbit, at the position of the maneuver node.</description></item><item><description>The z-axis points in the orbital normal direction of the original orbit,
at the position of the maneuver node.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		Node
```
#### Return 
```
ReferenceFrame```
### Orbit_static_ReferencePlaneNormal

```
<doc>
<summary>
The direction that is normal to the orbits reference plane,
in the given reference frame.
The reference plane is the plane from which the orbits inclination is measured.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Orbit_static_ReferencePlaneDirection

```
<doc>
<summary>
The direction from which the orbits longitude of ascending node is measured,
in the given reference frame.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Orbit_MeanAnomalyAtUT

```
<doc>
<summary>
The mean anomaly at the given time.
</summary>
<param name="ut">The universal time in seconds.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
ut		
```
#### Return 
```
```
### Orbit_RadiusAtTrueAnomaly

```
<doc>
<summary>
The orbital radius at the point in the orbit given by the true anomaly.
</summary>
<param name="trueAnomaly">The true anomaly.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
trueAnomaly		
```
#### Return 
```
```
### Orbit_TrueAnomalyAtRadius

```
<doc>
<summary>
The true anomaly at the given orbital radius.
</summary>
<param name="radius">The orbital radius in meters.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
radius		
```
#### Return 
```
```
### Orbit_TrueAnomalyAtUT

```
<doc>
<summary>
The true anomaly at the given time.
</summary>
<param name="ut">The universal time in seconds.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
ut		
```
#### Return 
```
```
### Orbit_UTAtTrueAnomaly

```
<doc>
<summary>
The universal time, in seconds, corresponding to the given true anomaly.
</summary>
<param name="trueAnomaly">True anomaly.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
trueAnomaly		
```
#### Return 
```
```
### Orbit_EccentricAnomalyAtUT

```
<doc>
<summary>
The eccentric anomaly at the given universal time.
</summary>
<param name="ut">The universal time, in seconds.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
ut		
```
#### Return 
```
```
### Orbit_OrbitalSpeedAt

```
<doc>
<summary>
The orbital speed at the given time, in meters per second.
</summary>
<param name="time">Time from now, in seconds.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
time		
```
#### Return 
```
```
### Orbit_RadiusAt

```
<doc>
<summary>
The orbital radius at the given time, in meters.
</summary>
<param name="ut">The universal time to measure the radius at.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
ut		
```
#### Return 
```
```
### Orbit_PositionAt

```
<doc>
<summary>
The position at a given time, in the specified reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="ut">The universal time to measure the position at.</param>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
ut		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Orbit_TimeOfClosestApproach

```
<doc>
<summary>
Estimates and returns the time at closest approach to a target vessel.
</summary>
<returns>The universal time at closest approach, in seconds.</returns>
<param name="target">Target vessel.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
#### Return 
```
```
### Orbit_DistanceAtClosestApproach

```
<doc>
<summary>
Estimates and returns the distance at closest approach to a target vessel, in meters.
</summary>
<param name="target">Target vessel.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
#### Return 
```
```
### Orbit_ListClosestApproaches

```
<doc>
<summary>
Returns the times at closest approach and corresponding distances, to a target vessel.
</summary>
<returns>
A list of two lists.
The first is a list of times at closest approach, as universal times in seconds.
The second is a list of corresponding distances at closest approach, in meters.
</returns>
<param name="target">Target vessel.</param>
<param name="orbits">The number of future orbits to search.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
```
orbits		
```
#### Return 
```
```
### Orbit_TrueAnomalyAtAN

```
<doc>
<summary>
The true anomaly of the ascending node with the given target vessel.
</summary>
<param name="target">Target vessel.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
#### Return 
```
```
### Orbit_TrueAnomalyAtDN

```
<doc>
<summary>
The true anomaly of the descending node with the given target vessel.
</summary>
<param name="target">Target vessel.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
#### Return 
```
```
### Orbit_RelativeInclination

```
<doc>
<summary>
Relative inclination of this orbit and the orbit of the given target vessel, in radians.
</summary>
<param name="target">Target vessel.</param>
</doc>
```
#### Parameters 
```
this		Orbit
```
```
target		Vessel
```
#### Return 
```
```
### Orbit_get_Body

```
<doc>
<summary>
The celestial body (e.g. planet or moon) around which the object is orbiting.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
CelestialBody```
### Orbit_get_Apoapsis

```
<doc>
<summary>
Gets the apoapsis of the orbit, in meters, from the center of mass
of the body being orbited.
</summary>
<remarks>
For the apoapsis altitude reported on the in-game map view,
use <see cref="M:SpaceCenter.Orbit.ApoapsisAltitude" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Periapsis

```
<doc>
<summary>
The periapsis of the orbit, in meters, from the center of mass
of the body being orbited.
</summary>
<remarks>
For the periapsis altitude reported on the in-game map view,
use <see cref="M:SpaceCenter.Orbit.PeriapsisAltitude" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_ApoapsisAltitude

```
<doc>
<summary>
The apoapsis of the orbit, in meters, above the sea level of the body being orbited.
</summary>
<remarks>
This is equal to <see cref="M:SpaceCenter.Orbit.Apoapsis" /> minus the equatorial radius of the body.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_PeriapsisAltitude

```
<doc>
<summary>
The periapsis of the orbit, in meters, above the sea level of the body being orbited.
</summary>
<remarks>
This is equal to <see cref="M:SpaceCenter.Orbit.Periapsis" /> minus the equatorial radius of the body.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_SemiMajorAxis

```
<doc>
<summary>
The semi-major axis of the orbit, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_SemiMinorAxis

```
<doc>
<summary>
The semi-minor axis of the orbit, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Radius

```
<doc>
<summary>
The current radius of the orbit, in meters. This is the distance between the center
of mass of the object in orbit, and the center of mass of the body around which it
is orbiting.
</summary>
<remarks>
This value will change over time if the orbit is elliptical.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Speed

```
<doc>
<summary>
The current orbital speed of the object in meters per second.
</summary>
<remarks>
This value will change over time if the orbit is elliptical.
</remarks>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Period

```
<doc>
<summary>
The orbital period, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_TimeToApoapsis

```
<doc>
<summary>
The time until the object reaches apoapsis, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_TimeToPeriapsis

```
<doc>
<summary>
The time until the object reaches periapsis, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Eccentricity

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Orbital_eccentricity">eccentricity</a>
of the orbit.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Inclination

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Orbital_inclination">inclination</a>
of the orbit,
in radians.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_LongitudeOfAscendingNode

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Longitude_of_the_ascending_node">longitude of
the ascending node</a>, in radians.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_ArgumentOfPeriapsis

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Argument_of_periapsis">argument of
periapsis</a>, in radians.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_MeanAnomalyAtEpoch

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Mean_anomaly">mean anomaly at epoch</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_Epoch

```
<doc>
<summary>
The time since the epoch (the point at which the
<a href="https://en.wikipedia.org/wiki/Mean_anomaly">mean anomaly at epoch</a>
was measured, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_MeanAnomaly

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Mean_anomaly">mean anomaly</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_EccentricAnomaly

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/Eccentric_anomaly">eccentric anomaly</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_TrueAnomaly

```
<doc>
<summary>
The <a href="https://en.wikipedia.org/wiki/True_anomaly">true anomaly</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_NextOrbit

```
<doc>
<summary>
If the object is going to change sphere of influence in the future, returns the new
orbit after the change. Otherwise returns <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
Orbit```
### Orbit_get_TimeToSOIChange

```
<doc>
<summary>
The time until the object changes sphere of influence, in seconds. Returns <c>NaN</c>
if the object is not going to change sphere of influence.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Orbit_get_OrbitalSpeed

```
<doc>
<summary>
The current orbital speed in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Orbit
```
#### Return 
```
```
### Antenna_Transmit

```
<doc>
<summary>
Transmit data.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_Cancel

```
<doc>
<summary>
Cancel current transmission of data.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_Part

```
<doc>
<summary>
The part object for this antenna.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
Part```
### Antenna_get_State

```
<doc>
<summary>
The current state of the antenna.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
AntennaState```
### Antenna_get_Deployable

```
<doc>
<summary>
Whether the antenna is deployable.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_Deployed

```
<doc>
<summary>
Whether the antenna is deployed.
</summary>
<remarks>
Fixed antennas are always deployed.
Returns an error if you try to deploy a fixed antenna.
</remarks>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_set_Deployed

```
<doc>
<summary>
Whether the antenna is deployed.
</summary>
<remarks>
Fixed antennas are always deployed.
Returns an error if you try to deploy a fixed antenna.
</remarks>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		
```
#### Return 
```
```
### Antenna_get_CanTransmit

```
<doc>
<summary>
Whether data can be transmitted by this antenna.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_AllowPartial

```
<doc>
<summary>
Whether partial data transmission is permitted.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_set_AllowPartial

```
<doc>
<summary>
Whether partial data transmission is permitted.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
```
value		
```
#### Return 
```
```
### Antenna_get_Power

```
<doc>
<summary>
The power of the antenna.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_Combinable

```
<doc>
<summary>
Whether the antenna can be combined with other antennae on the vessel
to boost the power.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_CombinableExponent

```
<doc>
<summary>
Exponent used to calculate the combined power of multiple antennae on a vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_PacketInterval

```
<doc>
<summary>
Interval between sending packets in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_PacketSize

```
<doc>
<summary>
Amount of data sent per packet in Mits.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### Antenna_get_PacketResourceCost

```
<doc>
<summary>
Units of electric charge consumed per packet sent.
</summary>
</doc>
```
#### Parameters 
```
this		Antenna
```
#### Return 
```
```
### CargoBay_get_Part

```
<doc>
<summary>
The part object for this cargo bay.
</summary>
</doc>
```
#### Parameters 
```
this		CargoBay
```
#### Return 
```
Part```
### CargoBay_get_State

```
<doc>
<summary>
The state of the cargo bay.
</summary>
</doc>
```
#### Parameters 
```
this		CargoBay
```
#### Return 
```
CargoBayState```
### CargoBay_get_Open

```
<doc>
<summary>
Whether the cargo bay is open.
</summary>
</doc>
```
#### Parameters 
```
this		CargoBay
```
#### Return 
```
```
### CargoBay_set_Open

```
<doc>
<summary>
Whether the cargo bay is open.
</summary>
</doc>
```
#### Parameters 
```
this		CargoBay
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_Part

```
<doc>
<summary>
The part object for this control surface.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
Part```
### ControlSurface_get_PitchEnabled

```
<doc>
<summary>
Whether the control surface has pitch control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_PitchEnabled

```
<doc>
<summary>
Whether the control surface has pitch control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_YawEnabled

```
<doc>
<summary>
Whether the control surface has yaw control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_YawEnabled

```
<doc>
<summary>
Whether the control surface has yaw control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_RollEnabled

```
<doc>
<summary>
Whether the control surface has roll control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_RollEnabled

```
<doc>
<summary>
Whether the control surface has roll control enabled.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_AuthorityLimiter

```
<doc>
<summary>
The authority limiter for the control surface, which controls how far the
control surface will move.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_AuthorityLimiter

```
<doc>
<summary>
The authority limiter for the control surface, which controls how far the
control surface will move.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_Inverted

```
<doc>
<summary>
Whether the control surface movement is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_Inverted

```
<doc>
<summary>
Whether the control surface movement is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_Deployed

```
<doc>
<summary>
Whether the control surface has been fully deployed.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_set_Deployed

```
<doc>
<summary>
Whether the control surface has been fully deployed.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
```
value		
```
#### Return 
```
```
### ControlSurface_get_SurfaceArea

```
<doc>
<summary>
Surface area of the control surface in <math>m^2</math>.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### ControlSurface_get_AvailableTorque

```
<doc>
<summary>
The available torque, in Newton meters, that can be produced by this control surface,
in the positive and negative pitch, roll and yaw axes of the vessel. These axes
correspond to the coordinate axes of the <see cref="M:SpaceCenter.Vessel.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		ControlSurface
```
#### Return 
```
```
### Decoupler_Decouple

```
<doc>
<summary>
Fires the decoupler. Returns the new vessel created when the decoupler fires.
Throws an exception if the decoupler has already fired.
</summary>
<remarks>
When called, the active vessel may change. It is therefore possible that,
after calling this function, the object(s) returned by previous call(s) to
<see cref="M:SpaceCenter.ActiveVessel" /> no longer refer to the active vessel.
</remarks>
</doc>
```
#### Parameters 
```
this		Decoupler
```
#### Return 
```
Vessel```
### Decoupler_get_Part

```
<doc>
<summary>
The part object for this decoupler.
</summary>
</doc>
```
#### Parameters 
```
this		Decoupler
```
#### Return 
```
Part```
### Decoupler_get_Decoupled

```
<doc>
<summary>
Whether the decoupler has fired.
</summary>
</doc>
```
#### Parameters 
```
this		Decoupler
```
#### Return 
```
```
### Decoupler_get_Staged

```
<doc>
<summary>
Whether the decoupler is enabled in the staging sequence.
</summary>
</doc>
```
#### Parameters 
```
this		Decoupler
```
#### Return 
```
```
### Decoupler_get_Impulse

```
<doc>
<summary>
The impulse that the decoupler imparts when it is fired, in Newton seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Decoupler
```
#### Return 
```
```
### DockingPort_Undock

```
<doc>
<summary>
Undocks the docking port and returns the new <see cref="T:SpaceCenter.Vessel" /> that is created.
This method can be called for either docking port in a docked pair.
Throws an exception if the docking port is not docked to anything.
</summary>
<remarks>
When called, the active vessel may change. It is therefore possible that,
after calling this function, the object(s) returned by previous call(s) to
<see cref="M:SpaceCenter.ActiveVessel" /> no longer refer to the active vessel.
</remarks>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
Vessel```
### DockingPort_Position

```
<doc>
<summary>
The position of the docking port, in the given reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		DockingPort
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### DockingPort_Direction

```
<doc>
<summary>
The direction that docking port points in, in the given reference frame.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		DockingPort
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### DockingPort_Rotation

```
<doc>
<summary>
The rotation of the docking port, in the given reference frame.
</summary>
<returns>The rotation as a quaternion of the form <math>(x, y, z, w)</math>.</returns>
<param name="referenceFrame">The reference frame that the returned
rotation is in.</param>
</doc>
```
#### Parameters 
```
this		DockingPort
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### DockingPort_get_Part

```
<doc>
<summary>
The part object for this docking port.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
Part```
### DockingPort_get_State

```
<doc>
<summary>
The current state of the docking port.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
DockingPortState```
### DockingPort_get_DockedPart

```
<doc>
<summary>
The part that this docking port is docked to. Returns <c>null</c> if this
docking port is not docked to anything.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
Part```
### DockingPort_get_ReengageDistance

```
<doc>
<summary>
The distance a docking port must move away when it undocks before it
becomes ready to dock with another port, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
```
### DockingPort_get_HasShield

```
<doc>
<summary>
Whether the docking port has a shield.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
```
### DockingPort_get_Shielded

```
<doc>
<summary>
The state of the docking ports shield, if it has one.

Returns <c>true</c> if the docking port has a shield, and the shield is
closed. Otherwise returns <c>false</c>. When set to <c>true</c>, the shield is
closed, and when set to <c>false</c> the shield is opened. If the docking
port does not have a shield, setting this attribute has no effect.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
```
### DockingPort_set_Shielded

```
<doc>
<summary>
The state of the docking ports shield, if it has one.

Returns <c>true</c> if the docking port has a shield, and the shield is
closed. Otherwise returns <c>false</c>. When set to <c>true</c>, the shield is
closed, and when set to <c>false</c> the shield is opened. If the docking
port does not have a shield, setting this attribute has no effect.
</summary>
</doc>
```
#### Parameters 
```
this		DockingPort
```
```
value		
```
#### Return 
```
```
### DockingPort_get_ReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to this docking port, and
oriented with the port.
<list type="bullet"><item><description>The origin is at the position of the docking port.
</description></item><item><description>The axes rotate with the docking port.</description></item><item><description>The x-axis points out to the right side of the docking port.
</description></item><item><description>The y-axis points in the direction the docking port is facing.
</description></item><item><description>The z-axis points out of the bottom off the docking port.
</description></item></list></summary>
<remarks>
This reference frame is not necessarily equivalent to the reference frame
for the part, returned by <see cref="M:SpaceCenter.Part.ReferenceFrame" />.
</remarks>
</doc>
```
#### Parameters 
```
this		DockingPort
```
#### Return 
```
ReferenceFrame```
### Engine_ToggleMode

```
<doc>
<summary>
Toggle the current engine mode.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_Part

```
<doc>
<summary>
The part object for this engine.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
Part```
### Engine_get_Active

```
<doc>
<summary>
Whether the engine is active. Setting this attribute may have no effect,
depending on <see cref="M:SpaceCenter.Engine.CanShutdown" /> and <see cref="M:SpaceCenter.Engine.CanRestart" />.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_Active

```
<doc>
<summary>
Whether the engine is active. Setting this attribute may have no effect,
depending on <see cref="M:SpaceCenter.Engine.CanShutdown" /> and <see cref="M:SpaceCenter.Engine.CanRestart" />.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_Thrust

```
<doc>
<summary>
The current amount of thrust being produced by the engine, in Newtons.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_AvailableThrust

```
<doc>
<summary>
The amount of thrust, in Newtons, that would be produced by the engine
when activated and with its throttle set to 100%.
Returns zero if the engine does not have any fuel.
Takes the engine's current <see cref="M:SpaceCenter.Engine.ThrustLimit" /> and atmospheric conditions
into account.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_MaxThrust

```
<doc>
<summary>
The amount of thrust, in Newtons, that would be produced by the engine
when activated and fueled, with its throttle and throttle limiter set to 100%.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_MaxVacuumThrust

```
<doc>
<summary>
The maximum amount of thrust that can be produced by the engine in a
vacuum, in Newtons. This is the amount of thrust produced by the engine
when activated, <see cref="M:SpaceCenter.Engine.ThrustLimit" /> is set to 100%, the main
vessel's throttle is set to 100% and the engine is in a vacuum.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_ThrustLimit

```
<doc>
<summary>
The thrust limiter of the engine. A value between 0 and 1. Setting this
attribute may have no effect, for example the thrust limit for a solid
rocket booster cannot be changed in flight.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_ThrustLimit

```
<doc>
<summary>
The thrust limiter of the engine. A value between 0 and 1. Setting this
attribute may have no effect, for example the thrust limit for a solid
rocket booster cannot be changed in flight.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_Thrusters

```
<doc>
<summary>
The components of the engine that generate thrust.
</summary>
<remarks>
For example, this corresponds to the rocket nozzel on a solid rocket booster,
or the individual nozzels on a RAPIER engine.
The overall thrust produced by the engine, as reported by <see cref="M:SpaceCenter.Engine.AvailableThrust" />,
<see cref="M:SpaceCenter.Engine.MaxThrust" /> and others, is the sum of the thrust generated by each thruster.
</remarks>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_SpecificImpulse

```
<doc>
<summary>
The current specific impulse of the engine, in seconds. Returns zero
if the engine is not active.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_VacuumSpecificImpulse

```
<doc>
<summary>
The vacuum specific impulse of the engine, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_KerbinSeaLevelSpecificImpulse

```
<doc>
<summary>
The specific impulse of the engine at sea level on Kerbin, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_PropellantNames

```
<doc>
<summary>
The names of the propellants that the engine consumes.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_Propellants

```
<doc>
<summary>
The propellants that the engine consumes.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_PropellantRatios

```
<doc>
<summary>
The ratio of resources that the engine consumes. A dictionary mapping resource names
to the ratio at which they are consumed by the engine.
</summary>
<remarks>
For example, if the ratios are 0.6 for LiquidFuel and 0.4 for Oxidizer, then for every
0.6 units of LiquidFuel that the engine burns, it will burn 0.4 units of Oxidizer.
</remarks>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_HasFuel

```
<doc>
<summary>
Whether the engine has any fuel available.
</summary>
<remarks>
The engine must be activated for this property to update correctly.
</remarks>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_Throttle

```
<doc>
<summary>
The current throttle setting for the engine. A value between 0 and 1.
This is not necessarily the same as the vessel's main throttle
setting, as some engines take time to adjust their throttle
(such as jet engines).
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_ThrottleLocked

```
<doc>
<summary>
Whether the <see cref="M:SpaceCenter.Control.Throttle" /> affects the engine. For example,
this is <c>true</c> for liquid fueled rockets, and <c>false</c> for solid rocket
boosters.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_CanRestart

```
<doc>
<summary>
Whether the engine can be restarted once shutdown. If the engine cannot be shutdown,
returns <c>false</c>. For example, this is <c>true</c> for liquid fueled rockets
and <c>false</c> for solid rocket boosters.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_CanShutdown

```
<doc>
<summary>
Whether the engine can be shutdown once activated. For example, this is
<c>true</c> for liquid fueled rockets and <c>false</c> for solid rocket boosters.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_HasModes

```
<doc>
<summary>
Whether the engine has multiple modes of operation.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_Mode

```
<doc>
<summary>
The name of the current engine mode.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_Mode

```
<doc>
<summary>
The name of the current engine mode.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_Modes

```
<doc>
<summary>
The available modes for the engine.
A dictionary mapping mode names to <see cref="T:SpaceCenter.Engine" /> objects.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_AutoModeSwitch

```
<doc>
<summary>
Whether the engine will automatically switch modes.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_AutoModeSwitch

```
<doc>
<summary>
Whether the engine will automatically switch modes.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_Gimballed

```
<doc>
<summary>
Whether the engine is gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_GimbalRange

```
<doc>
<summary>
The range over which the gimbal can move, in degrees.
Returns 0 if the engine is not gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_get_GimbalLocked

```
<doc>
<summary>
Whether the engines gimbal is locked in place. Setting this attribute has
no effect if the engine is not gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_GimbalLocked

```
<doc>
<summary>
Whether the engines gimbal is locked in place. Setting this attribute has
no effect if the engine is not gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_GimbalLimit

```
<doc>
<summary>
The gimbal limiter of the engine. A value between 0 and 1.
Returns 0 if the gimbal is locked.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Engine_set_GimbalLimit

```
<doc>
<summary>
The gimbal limiter of the engine. A value between 0 and 1.
Returns 0 if the gimbal is locked.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
```
value		
```
#### Return 
```
```
### Engine_get_AvailableTorque

```
<doc>
<summary>
The available torque, in Newton meters, that can be produced by this engine,
in the positive and negative pitch, roll and yaw axes of the vessel. These axes
correspond to the coordinate axes of the <see cref="M:SpaceCenter.Vessel.ReferenceFrame" />.
Returns zero if the engine is inactive, or not gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Engine
```
#### Return 
```
```
### Experiment_Run

```
<doc>
<summary>
Run the experiment.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_Transmit

```
<doc>
<summary>
Transmit all experimental data contained by this part.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_Dump

```
<doc>
<summary>
Dump the experimental data contained by the experiment.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_Reset

```
<doc>
<summary>
Reset the experiment.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Part

```
<doc>
<summary>
The part object for this experiment.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
Part```
### Experiment_get_Inoperable

```
<doc>
<summary>
Whether the experiment is inoperable.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Deployed

```
<doc>
<summary>
Whether the experiment has been deployed.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Rerunnable

```
<doc>
<summary>
Whether the experiment can be re-run.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_HasData

```
<doc>
<summary>
Whether the experiment contains data.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Data

```
<doc>
<summary>
The data contained in this experiment.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Available

```
<doc>
<summary>
Determines if the experiment is available given the current conditions.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_Biome

```
<doc>
<summary>
The name of the biome the experiment is currently in.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
```
### Experiment_get_ScienceSubject

```
<doc>
<summary>
Containing information on the corresponding specific science result for the current
conditions. Returns <c>null</c> if the experiment is unavailable.
</summary>
</doc>
```
#### Parameters 
```
this		Experiment
```
#### Return 
```
ScienceSubject```
### Fairing_Jettison

```
<doc>
<summary>
Jettison the fairing. Has no effect if it has already been jettisoned.
</summary>
</doc>
```
#### Parameters 
```
this		Fairing
```
#### Return 
```
```
### Fairing_get_Part

```
<doc>
<summary>
The part object for this fairing.
</summary>
</doc>
```
#### Parameters 
```
this		Fairing
```
#### Return 
```
Part```
### Fairing_get_Jettisoned

```
<doc>
<summary>
Whether the fairing has been jettisoned.
</summary>
</doc>
```
#### Parameters 
```
this		Fairing
```
#### Return 
```
```
### Force_Remove

```
<doc>
<summary>
Remove the force.
</summary>
</doc>
```
#### Parameters 
```
this		Force
```
#### Return 
```
```
### Force_get_Part

```
<doc>
<summary>
The part that this force is applied to.
</summary>
</doc>
```
#### Parameters 
```
this		Force
```
#### Return 
```
Part```
### Force_get_ForceVector

```
<doc>
<summary>
The force vector, in Newtons.
</summary>
<returns>A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Force
```
#### Return 
```
```
### Force_set_ForceVector

```
<doc>
<summary>
The force vector, in Newtons.
</summary>
<returns>A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</returns>
</doc>
```
#### Parameters 
```
this		Force
```
```
value		
```
#### Return 
```
```
### Force_get_Position

```
<doc>
<summary>
The position at which the force acts, in reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The position as a vector.</returns>
</doc>
```
#### Parameters 
```
this		Force
```
#### Return 
```
```
### Force_set_Position

```
<doc>
<summary>
The position at which the force acts, in reference frame <see cref="T:SpaceCenter.ReferenceFrame" />.
</summary>
<returns>The position as a vector.</returns>
</doc>
```
#### Parameters 
```
this		Force
```
```
value		
```
#### Return 
```
```
### Force_get_ReferenceFrame

```
<doc>
<summary>
The reference frame of the force vector and position.
</summary>
</doc>
```
#### Parameters 
```
this		Force
```
#### Return 
```
ReferenceFrame```
### Force_set_ReferenceFrame

```
<doc>
<summary>
The reference frame of the force vector and position.
</summary>
</doc>
```
#### Parameters 
```
this		Force
```
```
value		ReferenceFrame
```
#### Return 
```
```
### Intake_get_Part

```
<doc>
<summary>
The part object for this intake.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
#### Return 
```
Part```
### Intake_get_Open

```
<doc>
<summary>
Whether the intake is open.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
#### Return 
```
```
### Intake_set_Open

```
<doc>
<summary>
Whether the intake is open.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
```
value		
```
#### Return 
```
```
### Intake_get_Speed

```
<doc>
<summary>
Speed of the flow into the intake, in <math>m/s</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
#### Return 
```
```
### Intake_get_Flow

```
<doc>
<summary>
The rate of flow into the intake, in units of resource per second.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
#### Return 
```
```
### Intake_get_Area

```
<doc>
<summary>
The area of the intake's opening, in square meters.
</summary>
</doc>
```
#### Parameters 
```
this		Intake
```
#### Return 
```
```
### LaunchClamp_Release

```
<doc>
<summary>
Releases the docking clamp. Has no effect if the clamp has already been released.
</summary>
</doc>
```
#### Parameters 
```
this		LaunchClamp
```
#### Return 
```
```
### LaunchClamp_get_Part

```
<doc>
<summary>
The part object for this launch clamp.
</summary>
</doc>
```
#### Parameters 
```
this		LaunchClamp
```
#### Return 
```
Part```
### Leg_get_Part

```
<doc>
<summary>
The part object for this landing leg.
</summary>
</doc>
```
#### Parameters 
```
this		Leg
```
#### Return 
```
Part```
### Leg_get_State

```
<doc>
<summary>
The current state of the landing leg.
</summary>
</doc>
```
#### Parameters 
```
this		Leg
```
#### Return 
```
LegState```
### Leg_get_Deployable

```
<doc>
<summary>
Whether the leg is deployable.
</summary>
</doc>
```
#### Parameters 
```
this		Leg
```
#### Return 
```
```
### Leg_get_Deployed

```
<doc>
<summary>
Whether the landing leg is deployed.
</summary>
<remarks>
Fixed landing legs are always deployed.
Returns an error if you try to deploy fixed landing gear.
</remarks>
</doc>
```
#### Parameters 
```
this		Leg
```
#### Return 
```
```
### Leg_set_Deployed

```
<doc>
<summary>
Whether the landing leg is deployed.
</summary>
<remarks>
Fixed landing legs are always deployed.
Returns an error if you try to deploy fixed landing gear.
</remarks>
</doc>
```
#### Parameters 
```
this		Leg
```
```
value		
```
#### Return 
```
```
### Leg_get_IsGrounded

```
<doc>
<summary>
Returns whether the leg is touching the ground.
</summary>
</doc>
```
#### Parameters 
```
this		Leg
```
#### Return 
```
```
### Light_get_Part

```
<doc>
<summary>
The part object for this light.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
#### Return 
```
Part```
### Light_get_Active

```
<doc>
<summary>
Whether the light is switched on.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
#### Return 
```
```
### Light_set_Active

```
<doc>
<summary>
Whether the light is switched on.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
```
value		
```
#### Return 
```
```
### Light_get_Color

```
<doc>
<summary>
The color of the light, as an RGB triple.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
#### Return 
```
```
### Light_set_Color

```
<doc>
<summary>
The color of the light, as an RGB triple.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
```
value		
```
#### Return 
```
```
### Light_get_PowerUsage

```
<doc>
<summary>
The current power usage, in units of charge per second.
</summary>
</doc>
```
#### Parameters 
```
this		Light
```
#### Return 
```
```
### Module_HasField

```
<doc>
<summary>
Returns <c>true</c> if the module has a field with the given name.
</summary>
<param name="name">Name of the field.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_GetField

```
<doc>
<summary>
Returns the value of a field.
</summary>
<param name="name">Name of the field.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_SetFieldInt

```
<doc>
<summary>
Set the value of a field to the given integer number.
</summary>
<param name="name">Name of the field.</param>
<param name="value">Value to set.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
```
value		
```
#### Return 
```
```
### Module_SetFieldFloat

```
<doc>
<summary>
Set the value of a field to the given floating point number.
</summary>
<param name="name">Name of the field.</param>
<param name="value">Value to set.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
```
value		
```
#### Return 
```
```
### Module_SetFieldString

```
<doc>
<summary>
Set the value of a field to the given string.
</summary>
<param name="name">Name of the field.</param>
<param name="value">Value to set.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
```
value		
```
#### Return 
```
```
### Module_ResetField

```
<doc>
<summary>
Set the value of a field to its original value.
</summary>
<param name="name">Name of the field.</param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_HasEvent

```
<doc>
<summary><c>true</c> if the module has an event with the given name.
</summary>
<param name="name"></param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_TriggerEvent

```
<doc>
<summary>
Trigger the named event. Equivalent to clicking the button in the right-click menu
of the part.
</summary>
<param name="name"></param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_HasAction

```
<doc>
<summary><c>true</c> if the part has an action with the given name.
</summary>
<param name="name"></param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
#### Return 
```
```
### Module_SetAction

```
<doc>
<summary>
Set the value of an action with the given name.
</summary>
<param name="name"></param>
<param name="value"></param>
</doc>
```
#### Parameters 
```
this		Module
```
```
name		
```
```
value		
```
#### Return 
```
```
### Module_get_Name

```
<doc>
<summary>
Name of the PartModule. For example, "ModuleEngines".
</summary>
</doc>
```
#### Parameters 
```
this		Module
```
#### Return 
```
```
### Module_get_Part

```
<doc>
<summary>
The part that contains this module.
</summary>
</doc>
```
#### Parameters 
```
this		Module
```
#### Return 
```
Part```
### Module_get_Fields

```
<doc>
<summary>
The modules field names and their associated values, as a dictionary.
These are the values visible in the right-click menu of the part.
</summary>
</doc>
```
#### Parameters 
```
this		Module
```
#### Return 
```
```
### Module_get_Events

```
<doc>
<summary>
A list of the names of all of the modules events. Events are the clickable buttons
visible in the right-click menu of the part.
</summary>
</doc>
```
#### Parameters 
```
this		Module
```
#### Return 
```
```
### Module_get_Actions

```
<doc>
<summary>
A list of all the names of the modules actions. These are the parts actions that can
be assigned to action groups in the in-game editor.
</summary>
</doc>
```
#### Parameters 
```
this		Module
```
#### Return 
```
```
### Parachute_Deploy

```
<doc>
<summary>
Deploys the parachute. This has no effect if the parachute has already
been deployed.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_Arm

```
<doc>
<summary>
Deploys the parachute. This has no effect if the parachute has already
been armed or deployed. Only applicable to RealChutes parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_get_Part

```
<doc>
<summary>
The part object for this parachute.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
Part```
### Parachute_get_Deployed

```
<doc>
<summary>
Whether the parachute has been deployed.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_get_Armed

```
<doc>
<summary>
Whether the parachute has been armed or deployed. Only applicable to
RealChutes parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_get_State

```
<doc>
<summary>
The current state of the parachute.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
ParachuteState```
### Parachute_get_DeployAltitude

```
<doc>
<summary>
The altitude at which the parachute will full deploy, in meters.
Only applicable to stock parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_set_DeployAltitude

```
<doc>
<summary>
The altitude at which the parachute will full deploy, in meters.
Only applicable to stock parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
```
value		
```
#### Return 
```
```
### Parachute_get_DeployMinPressure

```
<doc>
<summary>
The minimum pressure at which the parachute will semi-deploy, in atmospheres.
Only applicable to stock parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
#### Return 
```
```
### Parachute_set_DeployMinPressure

```
<doc>
<summary>
The minimum pressure at which the parachute will semi-deploy, in atmospheres.
Only applicable to stock parachutes.
</summary>
</doc>
```
#### Parameters 
```
this		Parachute
```
```
value		
```
#### Return 
```
```
### Part_Position

```
<doc>
<summary>
The position of the part in the given reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
<remarks>
This is a fixed position in the part, defined by the parts model.
It s not necessarily the same as the parts center of mass.
Use <see cref="M:SpaceCenter.Part.CenterOfMass" /> to get the parts center of mass.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_CenterOfMass

```
<doc>
<summary>
The position of the parts center of mass in the given reference frame.
If the part is physicsless, this is equivalent to <see cref="M:SpaceCenter.Part.Position" />.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_BoundingBox

```
<doc>
<summary>
The axis-aligned bounding box of the part in the given reference frame.
</summary>
<returns>The positions of the minimum and maximum vertices of the box,
as position vectors.</returns>
<param name="referenceFrame">The reference frame that the returned
position vectors are in.</param>
<remarks>
This is computed from the collision mesh of the part.
If the part is not collidable, the box has zero volume and is centered on
the <see cref="M:SpaceCenter.Part.Position" /> of the part.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_Direction

```
<doc>
<summary>
The direction the part points in, in the given reference frame.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_Velocity

```
<doc>
<summary>
The linear velocity of the part in the given reference frame.
</summary>
<returns>The velocity as a vector. The vector points in the direction of travel,
and its magnitude is the speed of the body in meters per second.</returns>
<param name="referenceFrame">The reference frame that the returned
velocity vector is in.</param>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_Rotation

```
<doc>
<summary>
The rotation of the part, in the given reference frame.
</summary>
<returns>The rotation as a quaternion of the form <math>(x, y, z, w)</math>.</returns>
<param name="referenceFrame">The reference frame that the returned
rotation is in.</param>
</doc>
```
#### Parameters 
```
this		Part
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_AddForce

```
<doc>
<summary>
Exert a constant force on the part, acting at the given position.
</summary>
<returns>An object that can be used to remove or modify the force.</returns>
<param name="force">A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</param>
<param name="position">The position at which the force acts, as a vector.</param>
<param name="referenceFrame">The reference frame that the
force and position are in.</param>
</doc>
```
#### Parameters 
```
this		Part
```
```
force		
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
Force```
### Part_InstantaneousForce

```
<doc>
<summary>
Exert an instantaneous force on the part, acting at the given position.
</summary>
<param name="force">A vector pointing in the direction that the force acts,
with its magnitude equal to the strength of the force in Newtons.</param>
<param name="position">The position at which the force acts, as a vector.</param>
<param name="referenceFrame">The reference frame that the
force and position are in.</param>
<remarks>The force is applied instantaneously in a single physics update.</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
```
force		
```
```
position		
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Part_get_Name

```
<doc>
<summary>
Internal name of the part, as used in
<a href="https://wiki.kerbalspaceprogram.com/wiki/CFG_File_Documentation">part cfg files</a>.
For example "Mark1-2Pod".
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Title

```
<doc>
<summary>
Title of the part, as shown when the part is right clicked in-game. For example "Mk1-2 Command Pod".
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Tag

```
<doc>
<summary>
The name tag for the part. Can be set to a custom string using the
in-game user interface.
</summary>
<remarks>
This requires either the
<a href="https://github.com/krpc/NameTag/releases/latest">NameTag</a> or
<a href="https://forum.kerbalspaceprogram.com/index.php?/topic/61827-/">kOS</a>
mod to be installed.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_set_Tag

```
<doc>
<summary>
The name tag for the part. Can be set to a custom string using the
in-game user interface.
</summary>
<remarks>
This requires either the
<a href="https://github.com/krpc/NameTag/releases/latest">NameTag</a> or
<a href="https://forum.kerbalspaceprogram.com/index.php?/topic/61827-/">kOS</a>
mod to be installed.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
```
value		
```
#### Return 
```
```
### Part_get_Highlighted

```
<doc>
<summary>
Whether the part is highlighted.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_set_Highlighted

```
<doc>
<summary>
Whether the part is highlighted.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
```
value		
```
#### Return 
```
```
### Part_get_HighlightColor

```
<doc>
<summary>
The color used to highlight the part, as an RGB triple.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_set_HighlightColor

```
<doc>
<summary>
The color used to highlight the part, as an RGB triple.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
```
value		
```
#### Return 
```
```
### Part_get_Cost

```
<doc>
<summary>
The cost of the part, in units of funds.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Vessel

```
<doc>
<summary>
The vessel that contains this part.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Vessel```
### Part_get_Parent

```
<doc>
<summary>
The parts parent. Returns <c>null</c> if the part does not have a parent.
This, in combination with <see cref="M:SpaceCenter.Part.Children" />, can be used to traverse the vessels
parts tree.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Part```
### Part_get_Children

```
<doc>
<summary>
The parts children. Returns an empty list if the part has no children.
This, in combination with <see cref="M:SpaceCenter.Part.Parent" />, can be used to traverse the vessels
parts tree.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_AxiallyAttached

```
<doc>
<summary>
Whether the part is axially attached to its parent, i.e. on the top
or bottom of its parent. If the part has no parent, returns <c>false</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_RadiallyAttached

```
<doc>
<summary>
Whether the part is radially attached to its parent, i.e. on the side of its parent.
If the part has no parent, returns <c>false</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Stage

```
<doc>
<summary>
The stage in which this part will be activated. Returns -1 if the part is not
activated by staging.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_DecoupleStage

```
<doc>
<summary>
The stage in which this part will be decoupled. Returns -1 if the part is never
decoupled from the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Massless

```
<doc>
<summary>
Whether the part is
<a href="https://wiki.kerbalspaceprogram.com/wiki/Massless_part">massless</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Mass

```
<doc>
<summary>
The current mass of the part, including resources it contains, in kilograms.
Returns zero if the part is massless.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_DryMass

```
<doc>
<summary>
The mass of the part, not including any resources it contains, in kilograms.
Returns zero if the part is massless.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Shielded

```
<doc>
<summary>
Whether the part is shielded from the exterior of the vessel, for example by a fairing.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_DynamicPressure

```
<doc>
<summary>
The dynamic pressure acting on the part, in Pascals.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ImpactTolerance

```
<doc>
<summary>
The impact tolerance of the part, in meters per second.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Temperature

```
<doc>
<summary>
Temperature of the part, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_SkinTemperature

```
<doc>
<summary>
Temperature of the skin of the part, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_MaxTemperature

```
<doc>
<summary>
Maximum temperature that the part can survive, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_MaxSkinTemperature

```
<doc>
<summary>
Maximum temperature that the skin of the part can survive, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalMass

```
<doc>
<summary>
A measure of how much energy it takes to increase the internal temperature of the part,
in Joules per Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalSkinMass

```
<doc>
<summary>
A measure of how much energy it takes to increase the skin temperature of the part,
in Joules per Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalResourceMass

```
<doc>
<summary>
A measure of how much energy it takes to increase the temperature of the resources
contained in the part, in Joules per Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalInternalFlux

```
<doc>
<summary>
The rate at which heat energy is begin generated by the part.
For example, some engines generate heat by combusting fuel.
Measured in energy per unit time, or power, in Watts.
A positive value means the part is gaining heat energy, and negative means it is losing
heat energy.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalConductionFlux

```
<doc>
<summary>
The rate at which heat energy is conducting into or out of the part via contact with
other parts. Measured in energy per unit time, or power, in Watts.
A positive value means the part is gaining heat energy, and negative means it is
losing heat energy.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalConvectionFlux

```
<doc>
<summary>
The rate at which heat energy is convecting into or out of the part from the
surrounding atmosphere. Measured in energy per unit time, or power, in Watts.
A positive value means the part is gaining heat energy, and negative means it is
losing heat energy.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalRadiationFlux

```
<doc>
<summary>
The rate at which heat energy is radiating into or out of the part from the surrounding
environment. Measured in energy per unit time, or power, in Watts.
A positive value means the part is gaining heat energy, and negative means it is
losing heat energy.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ThermalSkinToInternalFlux

```
<doc>
<summary>
The rate at which heat energy is transferring between the part's skin and its internals.
Measured in energy per unit time, or power, in Watts.
A positive value means the part's internals are gaining heat energy,
and negative means its skin is gaining heat energy.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Resources

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Resources" /> object for the part.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Resources```
### Part_get_Crossfeed

```
<doc>
<summary>
Whether this part is crossfeed capable.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_IsFuelLine

```
<doc>
<summary>
Whether this part is a fuel line.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_FuelLinesFrom

```
<doc>
<summary>
The parts that are connected to this part via fuel lines, where the direction of the
fuel line is into this part.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_FuelLinesTo

```
<doc>
<summary>
The parts that are connected to this part via fuel lines, where the direction of the
fuel line is out of this part.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Modules

```
<doc>
<summary>
The modules for this part.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_Antenna

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Antenna" /> if the part is an antenna, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Antenna```
### Part_get_CargoBay

```
<doc>
<summary>
A <see cref="T:SpaceCenter.CargoBay" /> if the part is a cargo bay, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
CargoBay```
### Part_get_ControlSurface

```
<doc>
<summary>
A <see cref="T:SpaceCenter.ControlSurface" /> if the part is an aerodynamic control surface,
otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ControlSurface```
### Part_get_Decoupler

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Decoupler" /> if the part is a decoupler, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Decoupler```
### Part_get_DockingPort

```
<doc>
<summary>
A <see cref="T:SpaceCenter.DockingPort" /> if the part is a docking port, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
DockingPort```
### Part_get_Engine

```
<doc>
<summary>
An <see cref="T:SpaceCenter.Engine" /> if the part is an engine, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Engine```
### Part_get_Experiment

```
<doc>
<summary>
An <see cref="T:SpaceCenter.Experiment" /> if the part is a science experiment, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Experiment```
### Part_get_Fairing

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Fairing" /> if the part is a fairing, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Fairing```
### Part_get_Intake

```
<doc>
<summary>
An <see cref="T:SpaceCenter.Intake" /> if the part is an intake, otherwise <c>null</c>.
</summary>
<remarks>
This includes any part that generates thrust. This covers many different types
of engine, including liquid fuel rockets, solid rocket boosters and jet engines.
For RCS thrusters see <see cref="T:SpaceCenter.RCS" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Intake```
### Part_get_Leg

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Leg" /> if the part is a landing leg, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Leg```
### Part_get_LaunchClamp

```
<doc>
<summary>
A <see cref="T:SpaceCenter.LaunchClamp" /> if the part is a launch clamp, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
LaunchClamp```
### Part_get_Light

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Light" /> if the part is a light, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Light```
### Part_get_Parachute

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Parachute" /> if the part is a parachute, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Parachute```
### Part_get_Radiator

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Radiator" /> if the part is a radiator, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Radiator```
### Part_get_RCS

```
<doc>
<summary>
A <see cref="T:SpaceCenter.RCS" /> if the part is an RCS block/thruster, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
RCS```
### Part_get_ReactionWheel

```
<doc>
<summary>
A <see cref="T:SpaceCenter.ReactionWheel" /> if the part is a reaction wheel, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ReactionWheel```
### Part_get_ResourceConverter

```
<doc>
<summary>
A <see cref="T:SpaceCenter.ResourceConverter" /> if the part is a resource converter,
otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ResourceConverter```
### Part_get_ResourceHarvester

```
<doc>
<summary>
A <see cref="T:SpaceCenter.ResourceHarvester" /> if the part is a resource harvester,
otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ResourceHarvester```
### Part_get_Sensor

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Sensor" /> if the part is a sensor, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Sensor```
### Part_get_SolarPanel

```
<doc>
<summary>
A <see cref="T:SpaceCenter.SolarPanel" /> if the part is a solar panel, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
SolarPanel```
### Part_get_Wheel

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Wheel" /> if the part is a wheel, otherwise <c>null</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
Wheel```
### Part_get_MomentOfInertia

```
<doc>
<summary>
The moment of inertia of the part in <math>kg.m^2</math> around its center of mass
in the parts reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_InertiaTensor

```
<doc>
<summary>
The inertia tensor of the part in the parts reference frame
(<see cref="T:SpaceCenter.ReferenceFrame" />).
Returns the 3x3 matrix as a list of elements, in row-major order.
</summary>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
```
### Part_get_ReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to this part, and centered on a fixed
position within the part, defined by the parts model.
<list type="bullet"><item><description>The origin is at the position of the part, as returned by
<see cref="M:SpaceCenter.Part.Position" />.</description></item><item><description>The axes rotate with the part.</description></item><item><description>The x, y and z axis directions depend on the design of the part.
</description></item></list></summary>
<remarks>
For docking port parts, this reference frame is not necessarily equivalent to the
reference frame for the docking port, returned by
<see cref="M:SpaceCenter.DockingPort.ReferenceFrame" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ReferenceFrame```
### Part_get_CenterOfMassReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to this part, and centered on its
center of mass.
<list type="bullet"><item><description>The origin is at the center of mass of the part, as returned by
<see cref="M:SpaceCenter.Part.CenterOfMass" />.</description></item><item><description>The axes rotate with the part.</description></item><item><description>The x, y and z axis directions depend on the design of the part.
</description></item></list></summary>
<remarks>
For docking port parts, this reference frame is not necessarily equivalent to the
reference frame for the docking port, returned by
<see cref="M:SpaceCenter.DockingPort.ReferenceFrame" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Part
```
#### Return 
```
ReferenceFrame```
### Parts_WithName

```
<doc>
<summary>
A list of parts whose <see cref="M:SpaceCenter.Part.Name" /> is <paramref name="name" />.
</summary>
<param name="name"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
name		
```
#### Return 
```
```
### Parts_WithTitle

```
<doc>
<summary>
A list of all parts whose <see cref="M:SpaceCenter.Part.Title" /> is <paramref name="title" />.
</summary>
<param name="title"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
title		
```
#### Return 
```
```
### Parts_WithTag

```
<doc>
<summary>
A list of all parts whose <see cref="M:SpaceCenter.Part.Tag" /> is <paramref name="tag" />.
</summary>
<param name="tag"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
tag		
```
#### Return 
```
```
### Parts_WithModule

```
<doc>
<summary>
A list of all parts that contain a <see cref="T:SpaceCenter.Module" /> whose
<see cref="M:SpaceCenter.Module.Name" /> is <paramref name="moduleName" />.
</summary>
<param name="moduleName"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
moduleName		
```
#### Return 
```
```
### Parts_InStage

```
<doc>
<summary>
A list of all parts that are activated in the given <paramref name="stage" />.
</summary>
<param name="stage"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
stage		
```
#### Return 
```
```
### Parts_InDecoupleStage

```
<doc>
<summary>
A list of all parts that are decoupled in the given <paramref name="stage" />.
</summary>
<param name="stage"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
stage		
```
#### Return 
```
```
### Parts_ModulesWithName

```
<doc>
<summary>
A list of modules (combined across all parts in the vessel) whose
<see cref="M:SpaceCenter.Module.Name" /> is <paramref name="moduleName" />.
</summary>
<param name="moduleName"></param>
</doc>
```
#### Parameters 
```
this		Parts
```
```
moduleName		
```
#### Return 
```
```
### Parts_get_All

```
<doc>
<summary>
A list of all of the vessels parts.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Root

```
<doc>
<summary>
The vessels root part.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
Part```
### Parts_get_Controlling

```
<doc>
<summary>
The part from which the vessel is controlled.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
Part```
### Parts_set_Controlling

```
<doc>
<summary>
The part from which the vessel is controlled.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
```
value		Part
```
#### Return 
```
```
### Parts_get_Antennas

```
<doc>
<summary>
A list of all antennas in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_ControlSurfaces

```
<doc>
<summary>
A list of all control surfaces in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_CargoBays

```
<doc>
<summary>
A list of all cargo bays in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Decouplers

```
<doc>
<summary>
A list of all decouplers in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_DockingPorts

```
<doc>
<summary>
A list of all docking ports in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Engines

```
<doc>
<summary>
A list of all engines in the vessel.
</summary>
<remarks>
This includes any part that generates thrust. This covers many different types
of engine, including liquid fuel rockets, solid rocket boosters, jet engines and
RCS thrusters.
</remarks>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Experiments

```
<doc>
<summary>
A list of all science experiments in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Fairings

```
<doc>
<summary>
A list of all fairings in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Intakes

```
<doc>
<summary>
A list of all intakes in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Legs

```
<doc>
<summary>
A list of all landing legs attached to the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_LaunchClamps

```
<doc>
<summary>
A list of all launch clamps attached to the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Lights

```
<doc>
<summary>
A list of all lights in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Parachutes

```
<doc>
<summary>
A list of all parachutes in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Radiators

```
<doc>
<summary>
A list of all radiators in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_RCS

```
<doc>
<summary>
A list of all RCS blocks/thrusters in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_ReactionWheels

```
<doc>
<summary>
A list of all reaction wheels in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_ResourceConverters

```
<doc>
<summary>
A list of all resource converters in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_ResourceHarvesters

```
<doc>
<summary>
A list of all resource harvesters in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Sensors

```
<doc>
<summary>
A list of all sensors in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_SolarPanels

```
<doc>
<summary>
A list of all solar panels in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Parts_get_Wheels

```
<doc>
<summary>
A list of all wheels in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Parts
```
#### Return 
```
```
### Propellant_get_Name

```
<doc>
<summary>
The name of the propellant.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_CurrentAmount

```
<doc>
<summary>
The current amount of propellant.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_CurrentRequirement

```
<doc>
<summary>
The required amount of propellant.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_TotalResourceAvailable

```
<doc>
<summary>
The total amount of the underlying resource currently reachable given
resource flow rules.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_TotalResourceCapacity

```
<doc>
<summary>
The total vehicle capacity for the underlying propellant resource,
restricted by resource flow rules.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_IgnoreForIsp

```
<doc>
<summary>
If this propellant should be ignored when calculating required mass flow
given specific impulse.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_IgnoreForThrustCurve

```
<doc>
<summary>
If this propellant should be ignored for thrust curve calculations.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_DrawStackGauge

```
<doc>
<summary>
If this propellant has a stack gauge or not.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_IsDeprived

```
<doc>
<summary>
If this propellant is deprived.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### Propellant_get_Ratio

```
<doc>
<summary>
The propellant ratio.
</summary>
</doc>
```
#### Parameters 
```
this		Propellant
```
#### Return 
```
```
### RCS_get_Part

```
<doc>
<summary>
The part object for this RCS.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
Part```
### RCS_get_Active

```
<doc>
<summary>
Whether the RCS thrusters are active.
An RCS thruster is inactive if the RCS action group is disabled
(<see cref="M:SpaceCenter.Control.RCS" />), the RCS thruster itself is not enabled
(<see cref="M:SpaceCenter.RCS.Enabled" />) or it is covered by a fairing (<see cref="M:SpaceCenter.Part.Shielded" />).
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_Enabled

```
<doc>
<summary>
Whether the RCS thrusters are enabled.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_Enabled

```
<doc>
<summary>
Whether the RCS thrusters are enabled.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_PitchEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when pitch control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_PitchEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when pitch control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_YawEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when yaw control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_YawEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when yaw control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_RollEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when roll control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_RollEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when roll control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_ForwardEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when pitch control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_ForwardEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when pitch control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_UpEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when yaw control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_UpEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when yaw control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_RightEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when roll control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_set_RightEnabled

```
<doc>
<summary>
Whether the RCS thruster will fire when roll control input is given.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
```
value		
```
#### Return 
```
```
### RCS_get_AvailableTorque

```
<doc>
<summary>
The available torque, in Newton meters, that can be produced by this RCS,
in the positive and negative pitch, roll and yaw axes of the vessel. These axes
correspond to the coordinate axes of the <see cref="M:SpaceCenter.Vessel.ReferenceFrame" />.
Returns zero if RCS is disable.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_MaxThrust

```
<doc>
<summary>
The maximum amount of thrust that can be produced by the RCS thrusters when active,
in Newtons.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_MaxVacuumThrust

```
<doc>
<summary>
The maximum amount of thrust that can be produced by the RCS thrusters when active
in a vacuum, in Newtons.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_Thrusters

```
<doc>
<summary>
A list of thrusters, one of each nozzel in the RCS part.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_SpecificImpulse

```
<doc>
<summary>
The current specific impulse of the RCS, in seconds. Returns zero
if the RCS is not active.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_VacuumSpecificImpulse

```
<doc>
<summary>
The vacuum specific impulse of the RCS, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_KerbinSeaLevelSpecificImpulse

```
<doc>
<summary>
The specific impulse of the RCS at sea level on Kerbin, in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_Propellants

```
<doc>
<summary>
The names of resources that the RCS consumes.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_PropellantRatios

```
<doc>
<summary>
The ratios of resources that the RCS consumes. A dictionary mapping resource names
to the ratios at which they are consumed by the RCS.
</summary>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### RCS_get_HasFuel

```
<doc>
<summary>
Whether the RCS has fuel available.
</summary>
<remarks>
The RCS thruster must be activated for this property to update correctly.
</remarks>
</doc>
```
#### Parameters 
```
this		RCS
```
#### Return 
```
```
### Radiator_get_Part

```
<doc>
<summary>
The part object for this radiator.
</summary>
</doc>
```
#### Parameters 
```
this		Radiator
```
#### Return 
```
Part```
### Radiator_get_Deployable

```
<doc>
<summary>
Whether the radiator is deployable.
</summary>
</doc>
```
#### Parameters 
```
this		Radiator
```
#### Return 
```
```
### Radiator_get_Deployed

```
<doc>
<summary>
For a deployable radiator, <c>true</c> if the radiator is extended.
If the radiator is not deployable, this is always <c>true</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Radiator
```
#### Return 
```
```
### Radiator_set_Deployed

```
<doc>
<summary>
For a deployable radiator, <c>true</c> if the radiator is extended.
If the radiator is not deployable, this is always <c>true</c>.
</summary>
</doc>
```
#### Parameters 
```
this		Radiator
```
```
value		
```
#### Return 
```
```
### Radiator_get_State

```
<doc>
<summary>
The current state of the radiator.
</summary>
<remarks>
A fixed radiator is always <see cref="M:SpaceCenter.RadiatorState.Extended" />.
</remarks>
</doc>
```
#### Parameters 
```
this		Radiator
```
#### Return 
```
RadiatorState```
### ReactionWheel_get_Part

```
<doc>
<summary>
The part object for this reaction wheel.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
#### Return 
```
Part```
### ReactionWheel_get_Active

```
<doc>
<summary>
Whether the reaction wheel is active.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
#### Return 
```
```
### ReactionWheel_set_Active

```
<doc>
<summary>
Whether the reaction wheel is active.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
```
value		
```
#### Return 
```
```
### ReactionWheel_get_Broken

```
<doc>
<summary>
Whether the reaction wheel is broken.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
#### Return 
```
```
### ReactionWheel_get_AvailableTorque

```
<doc>
<summary>
The available torque, in Newton meters, that can be produced by this reaction wheel,
in the positive and negative pitch, roll and yaw axes of the vessel. These axes
correspond to the coordinate axes of the <see cref="M:SpaceCenter.Vessel.ReferenceFrame" />.
Returns zero if the reaction wheel is inactive or broken.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
#### Return 
```
```
### ReactionWheel_get_MaxTorque

```
<doc>
<summary>
The maximum torque, in Newton meters, that can be produced by this reaction wheel,
when it is active, in the positive and negative pitch, roll and yaw axes of the vessel.
These axes correspond to the coordinate axes of the <see cref="M:SpaceCenter.Vessel.ReferenceFrame" />.
</summary>
</doc>
```
#### Parameters 
```
this		ReactionWheel
```
#### Return 
```
```
### ResourceConverter_Active

```
<doc>
<summary>
True if the specified converter is active.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_Name

```
<doc>
<summary>
The name of the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_Start

```
<doc>
<summary>
Start the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_Stop

```
<doc>
<summary>
Stop the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_State

```
<doc>
<summary>
The state of the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
ResourceConverterState```
### ResourceConverter_StatusInfo

```
<doc>
<summary>
Status information for the specified converter.
This is the full status message shown in the in-game UI.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_Inputs

```
<doc>
<summary>
List of the names of resources consumed by the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_Outputs

```
<doc>
<summary>
List of the names of resources produced by the specified converter.
</summary>
<param name="index">Index of the converter.</param>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
```
index		
```
#### Return 
```
```
### ResourceConverter_get_Part

```
<doc>
<summary>
The part object for this converter.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
#### Return 
```
Part```
### ResourceConverter_get_Count

```
<doc>
<summary>
The number of converters in the part.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
#### Return 
```
```
### ResourceConverter_get_ThermalEfficiency

```
<doc>
<summary>
The thermal efficiency of the converter, as a percentage of its maximum.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
#### Return 
```
```
### ResourceConverter_get_CoreTemperature

```
<doc>
<summary>
The core temperature of the converter, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
#### Return 
```
```
### ResourceConverter_get_OptimumCoreTemperature

```
<doc>
<summary>
The core temperature at which the converter will operate with peak efficiency, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceConverter
```
#### Return 
```
```
### ResourceHarvester_get_Part

```
<doc>
<summary>
The part object for this harvester.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
Part```
### ResourceHarvester_get_State

```
<doc>
<summary>
The state of the harvester.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
ResourceHarvesterState```
### ResourceHarvester_get_Deployed

```
<doc>
<summary>
Whether the harvester is deployed.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ResourceHarvester_set_Deployed

```
<doc>
<summary>
Whether the harvester is deployed.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
```
value		
```
#### Return 
```
```
### ResourceHarvester_get_Active

```
<doc>
<summary>
Whether the harvester is actively drilling.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ResourceHarvester_set_Active

```
<doc>
<summary>
Whether the harvester is actively drilling.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
```
value		
```
#### Return 
```
```
### ResourceHarvester_get_ExtractionRate

```
<doc>
<summary>
The rate at which the drill is extracting ore, in units per second.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ResourceHarvester_get_ThermalEfficiency

```
<doc>
<summary>
The thermal efficiency of the drill, as a percentage of its maximum.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ResourceHarvester_get_CoreTemperature

```
<doc>
<summary>
The core temperature of the drill, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ResourceHarvester_get_OptimumCoreTemperature

```
<doc>
<summary>
The core temperature at which the drill will operate with peak efficiency, in Kelvin.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceHarvester
```
#### Return 
```
```
### ScienceData_get_DataAmount

```
<doc>
<summary>
Data amount.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceData
```
#### Return 
```
```
### ScienceData_get_ScienceValue

```
<doc>
<summary>
Science value.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceData
```
#### Return 
```
```
### ScienceData_get_TransmitValue

```
<doc>
<summary>
Transmit value.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceData
```
#### Return 
```
```
### ScienceSubject_get_Science

```
<doc>
<summary>
Amount of science already earned from this subject, not updated until after
transmission/recovery.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_ScienceCap

```
<doc>
<summary>
Total science allowable for this subject.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_IsComplete

```
<doc>
<summary>
Whether the experiment has been completed.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_DataScale

```
<doc>
<summary>
Multiply science value by this to determine data amount in mits.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_ScientificValue

```
<doc>
<summary>
Diminishing value multiplier for decreasing the science value returned from repeated
experiments.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_SubjectValue

```
<doc>
<summary>
Multiplier for specific Celestial Body/Experiment Situation combination.
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### ScienceSubject_get_Title

```
<doc>
<summary>
Title of science subject, displayed in science archives
</summary>
</doc>
```
#### Parameters 
```
this		ScienceSubject
```
#### Return 
```
```
### Sensor_get_Part

```
<doc>
<summary>
The part object for this sensor.
</summary>
</doc>
```
#### Parameters 
```
this		Sensor
```
#### Return 
```
Part```
### Sensor_get_Active

```
<doc>
<summary>
Whether the sensor is active.
</summary>
</doc>
```
#### Parameters 
```
this		Sensor
```
#### Return 
```
```
### Sensor_set_Active

```
<doc>
<summary>
Whether the sensor is active.
</summary>
</doc>
```
#### Parameters 
```
this		Sensor
```
```
value		
```
#### Return 
```
```
### Sensor_get_Value

```
<doc>
<summary>
The current value of the sensor.
</summary>
</doc>
```
#### Parameters 
```
this		Sensor
```
#### Return 
```
```
### SolarPanel_get_Part

```
<doc>
<summary>
The part object for this solar panel.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
Part```
### SolarPanel_get_Deployable

```
<doc>
<summary>
Whether the solar panel is deployable.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
```
### SolarPanel_get_Deployed

```
<doc>
<summary>
Whether the solar panel is extended.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
```
### SolarPanel_set_Deployed

```
<doc>
<summary>
Whether the solar panel is extended.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
```
value		
```
#### Return 
```
```
### SolarPanel_get_State

```
<doc>
<summary>
The current state of the solar panel.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
SolarPanelState```
### SolarPanel_get_EnergyFlow

```
<doc>
<summary>
The current amount of energy being generated by the solar panel, in
units of charge per second.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
```
### SolarPanel_get_SunExposure

```
<doc>
<summary>
The current amount of sunlight that is incident on the solar panel,
as a percentage. A value between 0 and 1.
</summary>
</doc>
```
#### Parameters 
```
this		SolarPanel
```
#### Return 
```
```
### Thruster_ThrustPosition

```
<doc>
<summary>
The position at which the thruster generates thrust, in the given reference frame.
For gimballed engines, this takes into account the current rotation of the gimbal.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Thruster
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Thruster_ThrustDirection

```
<doc>
<summary>
The direction of the force generated by the thruster, in the given reference frame.
This is opposite to the direction in which the thruster expels propellant.
For gimballed engines, this takes into account the current rotation of the gimbal.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		Thruster
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Thruster_InitialThrustPosition

```
<doc>
<summary>
The position at which the thruster generates thrust, when the engine is in its
initial position (no gimballing), in the given reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
<remarks>
This position can move when the gimbal rotates. This is because the thrust position and
gimbal position are not necessarily the same.
</remarks>
</doc>
```
#### Parameters 
```
this		Thruster
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Thruster_InitialThrustDirection

```
<doc>
<summary>
The direction of the force generated by the thruster, when the engine is in its
initial position (no gimballing), in the given reference frame.
This is opposite to the direction in which the thruster expels propellant.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		Thruster
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Thruster_GimbalPosition

```
<doc>
<summary>
Position around which the gimbal pivots.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Thruster
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Thruster_get_Part

```
<doc>
<summary>
The <see cref="T:SpaceCenter.Part" /> that contains this thruster.
</summary>
</doc>
```
#### Parameters 
```
this		Thruster
```
#### Return 
```
Part```
### Thruster_get_ThrustReferenceFrame

```
<doc>
<summary>
A reference frame that is fixed relative to the thruster and orientated with
its thrust direction (<see cref="M:SpaceCenter.Thruster.ThrustDirection" />).
For gimballed engines, this takes into account the current rotation of the gimbal.
<list type="bullet"><item><description>
The origin is at the position of thrust for this thruster
(<see cref="M:SpaceCenter.Thruster.ThrustPosition" />).</description></item><item><description>
The axes rotate with the thrust direction.
This is the direction in which the thruster expels propellant, including any gimballing.
</description></item><item><description>The y-axis points along the thrust direction.</description></item><item><description>The x-axis and z-axis are perpendicular to the thrust direction.
</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		Thruster
```
#### Return 
```
ReferenceFrame```
### Thruster_get_Gimballed

```
<doc>
<summary>
Whether the thruster is gimballed.
</summary>
</doc>
```
#### Parameters 
```
this		Thruster
```
#### Return 
```
```
### Thruster_get_GimbalAngle

```
<doc>
<summary>
The current gimbal angle in the pitch, roll and yaw axes, in degrees.
</summary>
</doc>
```
#### Parameters 
```
this		Thruster
```
#### Return 
```
```
### Wheel_get_Part

```
<doc>
<summary>
The part object for this wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
Part```
### Wheel_get_State

```
<doc>
<summary>
The current state of the wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
WheelState```
### Wheel_get_Radius

```
<doc>
<summary>
Radius of the wheel, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Grounded

```
<doc>
<summary>
Whether the wheel is touching the ground.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_HasBrakes

```
<doc>
<summary>
Whether the wheel has brakes.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Brakes

```
<doc>
<summary>
The braking force, as a percentage of maximum, when the brakes are applied.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_Brakes

```
<doc>
<summary>
The braking force, as a percentage of maximum, when the brakes are applied.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_AutoFrictionControl

```
<doc>
<summary>
Whether automatic friction control is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_AutoFrictionControl

```
<doc>
<summary>
Whether automatic friction control is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_ManualFrictionControl

```
<doc>
<summary>
Manual friction control value. Only has an effect if automatic friction control is disabled.
A value between 0 and 5 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_ManualFrictionControl

```
<doc>
<summary>
Manual friction control value. Only has an effect if automatic friction control is disabled.
A value between 0 and 5 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_Deployable

```
<doc>
<summary>
Whether the wheel is deployable.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Deployed

```
<doc>
<summary>
Whether the wheel is deployed.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_Deployed

```
<doc>
<summary>
Whether the wheel is deployed.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_Powered

```
<doc>
<summary>
Whether the wheel is powered by a motor.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_MotorEnabled

```
<doc>
<summary>
Whether the motor is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_MotorEnabled

```
<doc>
<summary>
Whether the motor is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_MotorInverted

```
<doc>
<summary>
Whether the direction of the motor is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_MotorInverted

```
<doc>
<summary>
Whether the direction of the motor is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_MotorState

```
<doc>
<summary>
Whether the direction of the motor is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
MotorState```
### Wheel_get_MotorOutput

```
<doc>
<summary>
The output of the motor. This is the torque currently being generated, in Newton meters.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_TractionControlEnabled

```
<doc>
<summary>
Whether automatic traction control is enabled.
A wheel only has traction control if it is powered.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_TractionControlEnabled

```
<doc>
<summary>
Whether automatic traction control is enabled.
A wheel only has traction control if it is powered.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_TractionControl

```
<doc>
<summary>
Setting for the traction control.
Only takes effect if the wheel has automatic traction control enabled.
A value between 0 and 5 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_TractionControl

```
<doc>
<summary>
Setting for the traction control.
Only takes effect if the wheel has automatic traction control enabled.
A value between 0 and 5 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_DriveLimiter

```
<doc>
<summary>
Manual setting for the motor limiter.
Only takes effect if the wheel has automatic traction control disabled.
A value between 0 and 100 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_DriveLimiter

```
<doc>
<summary>
Manual setting for the motor limiter.
Only takes effect if the wheel has automatic traction control disabled.
A value between 0 and 100 inclusive.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_Steerable

```
<doc>
<summary>
Whether the wheel has steering.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_SteeringEnabled

```
<doc>
<summary>
Whether the wheel steering is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_SteeringEnabled

```
<doc>
<summary>
Whether the wheel steering is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_SteeringInverted

```
<doc>
<summary>
Whether the wheel steering is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_set_SteeringInverted

```
<doc>
<summary>
Whether the wheel steering is inverted.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
```
value		
```
#### Return 
```
```
### Wheel_get_HasSuspension

```
<doc>
<summary>
Whether the wheel has suspension.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_SuspensionSpringStrength

```
<doc>
<summary>
Suspension spring strength, as set in the editor.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_SuspensionDamperStrength

```
<doc>
<summary>
Suspension damper strength, as set in the editor.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Broken

```
<doc>
<summary>
Whether the wheel is broken.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Repairable

```
<doc>
<summary>
Whether the wheel is repairable.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Stress

```
<doc>
<summary>
Current stress on the wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_StressTolerance

```
<doc>
<summary>
Stress tolerance of the wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_StressPercentage

```
<doc>
<summary>
Current stress on the wheel as a percentage of its stress tolerance.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Deflection

```
<doc>
<summary>
Current deflection of the wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### Wheel_get_Slip

```
<doc>
<summary>
Current slip of the wheel.
</summary>
</doc>
```
#### Parameters 
```
this		Wheel
```
#### Return 
```
```
### ReferenceFrame_static_CreateRelative

```
<doc>
<summary>
Create a relative reference frame. This is a custom reference frame
whose components offset the components of a parent reference frame.
</summary>
<param name="referenceFrame">The parent reference frame on which to
base this reference frame.</param>
<param name="position">The offset of the position of the origin,
as a position vector. Defaults to <math>(0, 0, 0)</math></param>
<param name="rotation">The rotation to apply to the parent frames rotation,
as a quaternion of the form <math>(x, y, z, w)</math>.
Defaults to <math>(0, 0, 0, 1)</math> (i.e. no rotation)</param>
<param name="velocity">The linear velocity to offset the parent frame by,
as a vector pointing in the direction of travel, whose magnitude is the speed in
meters per second. Defaults to <math>(0, 0, 0)</math>.</param>
<param name="angularVelocity">The angular velocity to offset the parent frame by,
as a vector. This vector points in the direction of the axis of rotation,
and its magnitude is the speed of the rotation in radians per second.
Defaults to <math>(0, 0, 0)</math>.</param>
</doc>
```
#### Parameters 
```
referenceFrame		ReferenceFrame
```
```
position		
```
```
rotation		
```
```
velocity		
```
```
angularVelocity		
```
#### Return 
```
ReferenceFrame```
### ReferenceFrame_static_CreateHybrid

```
<doc>
<summary>
Create a hybrid reference frame. This is a custom reference frame
whose components inherited from other reference frames.
</summary>
<param name="position">The reference frame providing the position of the origin.</param>
<param name="rotation">The reference frame providing the rotation of the frame.</param>
<param name="velocity">The reference frame providing the linear velocity of the frame.
</param>
<param name="angularVelocity">The reference frame providing the angular velocity
of the frame.</param>
<remarks>
The <paramref name="position" /> reference frame is required but all other
reference frames are optional. If omitted, they are set to the
<paramref name="position" /> reference frame.
</remarks>
</doc>
```
#### Parameters 
```
position		ReferenceFrame
```
```
rotation		ReferenceFrame
```
```
velocity		ReferenceFrame
```
```
angularVelocity		ReferenceFrame
```
#### Return 
```
ReferenceFrame```
### Resource_get_Name

```
<doc>
<summary>
The name of the resource.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
```
### Resource_get_Part

```
<doc>
<summary>
The part containing the resource.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
Part```
### Resource_get_Max

```
<doc>
<summary>
The total amount of the resource that can be stored in the part.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
```
### Resource_get_Amount

```
<doc>
<summary>
The amount of the resource that is currently stored in the part.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
```
### Resource_get_Density

```
<doc>
<summary>
The density of the resource, in <math>kg/l</math>.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
```
### Resource_get_FlowMode

```
<doc>
<summary>
The flow mode of the resource.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
ResourceFlowMode```
### Resource_get_Enabled

```
<doc>
<summary>
Whether use of this resource is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
#### Return 
```
```
### Resource_set_Enabled

```
<doc>
<summary>
Whether use of this resource is enabled.
</summary>
</doc>
```
#### Parameters 
```
this		Resource
```
```
value		
```
#### Return 
```
```
### ResourceTransfer_static_Start

```
<doc>
<summary>
Start transferring a resource transfer between a pair of parts. The transfer will move
at most <paramref name="maxAmount" /> units of the resource, depending on how much of
the resource is available in the source part and how much storage is available in the
destination part.
Use <see cref="M:SpaceCenter.ResourceTransfer.Complete" /> to check if the transfer is complete.
Use <see cref="M:SpaceCenter.ResourceTransfer.Amount" /> to see how much of the resource has been transferred.
</summary>
<param name="fromPart">The part to transfer to.</param>
<param name="toPart">The part to transfer from.</param>
<param name="resource">The name of the resource to transfer.</param>
<param name="maxAmount">The maximum amount of resource to transfer.</param>
</doc>
```
#### Parameters 
```
fromPart		Part
```
```
toPart		Part
```
```
resource		
```
```
maxAmount		
```
#### Return 
```
ResourceTransfer```
### ResourceTransfer_get_Complete

```
<doc>
<summary>
Whether the transfer has completed.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceTransfer
```
#### Return 
```
```
### ResourceTransfer_get_Amount

```
<doc>
<summary>
The amount of the resource that has been transferred.
</summary>
</doc>
```
#### Parameters 
```
this		ResourceTransfer
```
#### Return 
```
```
### Resources_WithResource

```
<doc>
<summary>
All the individual resources with the given name that can be stored.
</summary>
</doc>
```
#### Parameters 
```
this		Resources
```
```
name		
```
#### Return 
```
```
### Resources_HasResource

```
<doc>
<summary>
Check whether the named resource can be stored.
</summary>
<param name="name">The name of the resource.</param>
</doc>
```
#### Parameters 
```
this		Resources
```
```
name		
```
#### Return 
```
```
### Resources_Max

```
<doc>
<summary>
Returns the amount of a resource that can be stored.
</summary>
<param name="name">The name of the resource.</param>
</doc>
```
#### Parameters 
```
this		Resources
```
```
name		
```
#### Return 
```
```
### Resources_Amount

```
<doc>
<summary>
Returns the amount of a resource that is currently stored.
</summary>
<param name="name">The name of the resource.</param>
</doc>
```
#### Parameters 
```
this		Resources
```
```
name		
```
#### Return 
```
```
### Resources_static_Density

```
<doc>
<summary>
Returns the density of a resource, in <math>kg/l</math>.
</summary>
<param name="name">The name of the resource.</param>
</doc>
```
#### Parameters 
```
name		
```
#### Return 
```
```
### Resources_static_FlowMode

```
<doc>
<summary>
Returns the flow mode of a resource.
</summary>
<param name="name">The name of the resource.</param>
</doc>
```
#### Parameters 
```
name		
```
#### Return 
```
ResourceFlowMode```
### Resources_get_All

```
<doc>
<summary>
All the individual resources that can be stored.
</summary>
</doc>
```
#### Parameters 
```
this		Resources
```
#### Return 
```
```
### Resources_get_Names

```
<doc>
<summary>
A list of resource names that can be stored.
</summary>
</doc>
```
#### Parameters 
```
this		Resources
```
#### Return 
```
```
### Resources_get_Enabled

```
<doc>
<summary>
Whether use of all the resources are enabled.
</summary>
<remarks>
This is <c>true</c> if all of the resources are enabled.
If any of the resources are not enabled, this is <c>false</c>.
</remarks>
</doc>
```
#### Parameters 
```
this		Resources
```
#### Return 
```
```
### Resources_set_Enabled

```
<doc>
<summary>
Whether use of all the resources are enabled.
</summary>
<remarks>
This is <c>true</c> if all of the resources are enabled.
If any of the resources are not enabled, this is <c>false</c>.
</remarks>
</doc>
```
#### Parameters 
```
this		Resources
```
```
value		
```
#### Return 
```
```
### Vessel_Recover

```
<doc>
<summary>
Recover the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_Flight

```
<doc>
<summary>
Returns a <see cref="T:SpaceCenter.Flight" /> object that can be used to get flight
telemetry for the vessel, in the specified reference frame.
</summary>
<param name="referenceFrame">
Reference frame. Defaults to the vessel's surface reference frame
(<see cref="M:SpaceCenter.Vessel.SurfaceReferenceFrame" />).
</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
Flight```
### Vessel_ResourcesInDecoupleStage

```
<doc>
<summary>
Returns a <see cref="T:SpaceCenter.Resources" /> object, that can used to get
information about resources stored in a given <paramref name="stage" />.
</summary>
<param name="stage">Get resources for parts that are decoupled in this stage.</param>
<param name="cumulative">When <c>false</c>, returns the resources for parts
decoupled in just the given stage. When <c>true</c> returns the resources decoupled in
the given stage and all subsequent stages combined.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
stage		
```
```
cumulative		
```
#### Return 
```
Resources```
### Vessel_Position

```
<doc>
<summary>
The position of the center of mass of the vessel, in the given reference frame.
</summary>
<returns>The position as a vector.</returns>
<param name="referenceFrame">The reference frame that the returned
position vector is in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_BoundingBox

```
<doc>
<summary>
The axis-aligned bounding box of the vessel in the given reference frame.
</summary>
<returns>The positions of the minimum and maximum vertices of the box,
as position vectors.</returns>
<param name="referenceFrame">The reference frame that the returned
position vectors are in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_Velocity

```
<doc>
<summary>
The velocity of the center of mass of the vessel, in the given reference frame.
</summary>
<returns>The velocity as a vector. The vector points in the direction of travel,
and its magnitude is the speed of the body in meters per second.</returns>
<param name="referenceFrame">The reference frame that the returned
velocity vector is in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_Rotation

```
<doc>
<summary>
The rotation of the vessel, in the given reference frame.
</summary>
<returns>The rotation as a quaternion of the form <math>(x, y, z, w)</math>.</returns>
<param name="referenceFrame">The reference frame that the returned
rotation is in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_Direction

```
<doc>
<summary>
The direction in which the vessel is pointing, in the given reference frame.
</summary>
<returns>The direction as a unit vector.</returns>
<param name="referenceFrame">The reference frame that the returned
direction is in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_AngularVelocity

```
<doc>
<summary>
The angular velocity of the vessel, in the given reference frame.
</summary>
<returns>The angular velocity as a vector. The magnitude of the vector is the rotational
speed of the vessel, in radians per second. The direction of the vector indicates the
axis of rotation, using the right-hand rule.</returns>
<param name="referenceFrame">The reference frame the returned
angular velocity is in.</param>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
referenceFrame		ReferenceFrame
```
#### Return 
```
```
### Vessel_get_Name

```
<doc>
<summary>
The name of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_set_Name

```
<doc>
<summary>
The name of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
value		
```
#### Return 
```
```
### Vessel_get_Type

```
<doc>
<summary>
The type of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
VesselType```
### Vessel_set_Type

```
<doc>
<summary>
The type of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
```
value		VesselType
```
#### Return 
```
```
### Vessel_get_Situation

```
<doc>
<summary>
The situation the vessel is in.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
VesselSituation```
### Vessel_get_Recoverable

```
<doc>
<summary>
Whether the vessel is recoverable.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_MET

```
<doc>
<summary>
The mission elapsed time in seconds.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_Biome

```
<doc>
<summary>
The name of the biome the vessel is currently in.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_Orbit

```
<doc>
<summary>
The current orbit of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
Orbit```
### Vessel_get_Control

```
<doc>
<summary>
Returns a <see cref="T:SpaceCenter.Control" /> object that can be used to manipulate
the vessel's control inputs. For example, its pitch/yaw/roll controls,
RCS and thrust.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
Control```
### Vessel_get_Comms

```
<doc>
<summary>
Returns a <see cref="T:SpaceCenter.Comms" /> object that can be used to interact
with CommNet for this vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
Comms```
### Vessel_get_AutoPilot

```
<doc>
<summary>
An <see cref="T:SpaceCenter.AutoPilot" /> object, that can be used to perform
simple auto-piloting of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
AutoPilot```
### Vessel_get_CrewCapacity

```
<doc>
<summary>
The number of crew that can occupy the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_CrewCount

```
<doc>
<summary>
The number of crew that are occupying the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_Crew

```
<doc>
<summary>
The crew in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_Resources

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Resources" /> object, that can used to get information
about resources stored in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
Resources```
### Vessel_get_Parts

```
<doc>
<summary>
A <see cref="T:SpaceCenter.Parts" /> object, that can used to interact with the parts that make up this vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
Parts```
### Vessel_get_Mass

```
<doc>
<summary>
The total mass of the vessel, including resources, in kg.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_DryMass

```
<doc>
<summary>
The total mass of the vessel, excluding resources, in kg.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_Thrust

```
<doc>
<summary>
The total thrust currently being produced by the vessel's engines, in
Newtons. This is computed by summing <see cref="M:SpaceCenter.Engine.Thrust" /> for
every engine in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableThrust

```
<doc>
<summary>
Gets the total available thrust that can be produced by the vessel's
active engines, in Newtons. This is computed by summing
<see cref="M:SpaceCenter.Engine.AvailableThrust" /> for every active engine in the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_MaxThrust

```
<doc>
<summary>
The total maximum thrust that can be produced by the vessel's active
engines, in Newtons. This is computed by summing
<see cref="M:SpaceCenter.Engine.MaxThrust" /> for every active engine.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_MaxVacuumThrust

```
<doc>
<summary>
The total maximum thrust that can be produced by the vessel's active
engines when the vessel is in a vacuum, in Newtons. This is computed by
summing <see cref="M:SpaceCenter.Engine.MaxVacuumThrust" /> for every active engine.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_SpecificImpulse

```
<doc>
<summary>
The combined specific impulse of all active engines, in seconds. This is computed using the formula
<a href="https://wiki.kerbalspaceprogram.com/wiki/Specific_impulse#Multiple_engines">described here</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_VacuumSpecificImpulse

```
<doc>
<summary>
The combined vacuum specific impulse of all active engines, in seconds. This is computed using the formula
<a href="https://wiki.kerbalspaceprogram.com/wiki/Specific_impulse#Multiple_engines">described here</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_KerbinSeaLevelSpecificImpulse

```
<doc>
<summary>
The combined specific impulse of all active engines at sea level on Kerbin, in seconds.
This is computed using the formula
<a href="https://wiki.kerbalspaceprogram.com/wiki/Specific_impulse#Multiple_engines">described here</a>.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_MomentOfInertia

```
<doc>
<summary>
The moment of inertia of the vessel around its center of mass in <math>kg.m^2</math>.
The inertia values in the returned 3-tuple are around the
pitch, roll and yaw directions respectively.
This corresponds to the vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_InertiaTensor

```
<doc>
<summary>
The inertia tensor of the vessel around its center of mass,
in the vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
Returns the 3x3 matrix as a list of elements, in row-major order.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableTorque

```
<doc>
<summary>
The maximum torque that the vessel generates. Includes contributions from
reaction wheels, RCS, gimballed engines and aerodynamic control surfaces.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableReactionWheelTorque

```
<doc>
<summary>
The maximum torque that the currently active and powered reaction wheels can generate.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableRCSTorque

```
<doc>
<summary>
The maximum torque that the currently active RCS thrusters can generate.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableEngineTorque

```
<doc>
<summary>
The maximum torque that the currently active and gimballed engines can generate.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableControlSurfaceTorque

```
<doc>
<summary>
The maximum torque that the aerodynamic control surfaces can generate.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_AvailableOtherTorque

```
<doc>
<summary>
The maximum torque that parts (excluding reaction wheels, gimballed engines,
RCS and control surfaces) can generate.
Returns the torques in <math>N.m</math> around each of the coordinate axes of the
vessels reference frame (<see cref="T:SpaceCenter.ReferenceFrame" />).
These axes are equivalent to the pitch, roll and yaw axes of the vessel.
</summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
```
### Vessel_get_ReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the vessel,
and orientated with the vessel.
<list type="bullet"><item><description>The origin is at the center of mass of the vessel.</description></item><item><description>The axes rotate with the vessel.</description></item><item><description>The x-axis points out to the right of the vessel.</description></item><item><description>The y-axis points in the forward direction of the vessel.</description></item><item><description>The z-axis points out of the bottom off the vessel.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
ReferenceFrame```
### Vessel_get_OrbitalReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the vessel,
and orientated with the vessels orbital prograde/normal/radial directions.
<list type="bullet"><item><description>The origin is at the center of mass of the vessel.</description></item><item><description>The axes rotate with the orbital prograde/normal/radial directions.</description></item><item><description>The x-axis points in the orbital anti-radial direction.</description></item><item><description>The y-axis points in the orbital prograde direction.</description></item><item><description>The z-axis points in the orbital normal direction.</description></item></list></summary>
<remarks>
Be careful not to confuse this with 'orbit' mode on the navball.
</remarks>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
ReferenceFrame```
### Vessel_get_SurfaceReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the vessel,
and orientated with the surface of the body being orbited.
<list type="bullet"><item><description>The origin is at the center of mass of the vessel.</description></item><item><description>The axes rotate with the north and up directions on the surface of the body.</description></item><item><description>The x-axis points in the <a href="https://en.wikipedia.org/wiki/Zenith">zenith</a>
direction (upwards, normal to the body being orbited, from the center of the body towards the center of
mass of the vessel).</description></item><item><description>The y-axis points northwards towards the
<a href="https://en.wikipedia.org/wiki/Horizon">astronomical horizon</a> (north, and tangential to the
surface of the body -- the direction in which a compass would point when on the surface).</description></item><item><description>The z-axis points eastwards towards the
<a href="https://en.wikipedia.org/wiki/Horizon">astronomical horizon</a> (east, and tangential to the
surface of the body -- east on a compass when on the surface).</description></item></list></summary>
<remarks>
Be careful not to confuse this with 'surface' mode on the navball.
</remarks>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
ReferenceFrame```
### Vessel_get_SurfaceVelocityReferenceFrame

```
<doc>
<summary>
The reference frame that is fixed relative to the vessel,
and orientated with the velocity vector of the vessel relative
to the surface of the body being orbited.
<list type="bullet"><item><description>The origin is at the center of mass of the vessel.</description></item><item><description>The axes rotate with the vessel's velocity vector.</description></item><item><description>The y-axis points in the direction of the vessel's velocity vector,
relative to the surface of the body being orbited.</description></item><item><description>The z-axis is in the plane of the
<a href="https://en.wikipedia.org/wiki/Horizon">astronomical horizon</a>.</description></item><item><description>The x-axis is orthogonal to the other two axes.</description></item></list></summary>
</doc>
```
#### Parameters 
```
this		Vessel
```
#### Return 
```
ReferenceFrame```
### Waypoint_Remove

```
<doc>
<summary>
Removes the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_Body

```
<doc>
<summary>
The celestial body the waypoint is attached to.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
CelestialBody```
### Waypoint_set_Body

```
<doc>
<summary>
The celestial body the waypoint is attached to.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		CelestialBody
```
#### Return 
```
```
### Waypoint_get_Name

```
<doc>
<summary>
The name of the waypoint as it appears on the map and the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_Name

```
<doc>
<summary>
The name of the waypoint as it appears on the map and the contract.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_Color

```
<doc>
<summary>
The seed of the icon color. See <see cref="M:SpaceCenter.WaypointManager.Colors" /> for example colors.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_Color

```
<doc>
<summary>
The seed of the icon color. See <see cref="M:SpaceCenter.WaypointManager.Colors" /> for example colors.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_Icon

```
<doc>
<summary>
The icon of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_Icon

```
<doc>
<summary>
The icon of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_Latitude

```
<doc>
<summary>
The latitude of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_Latitude

```
<doc>
<summary>
The latitude of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_Longitude

```
<doc>
<summary>
The longitude of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_Longitude

```
<doc>
<summary>
The longitude of the waypoint.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_MeanAltitude

```
<doc>
<summary>
The altitude of the waypoint above sea level, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_MeanAltitude

```
<doc>
<summary>
The altitude of the waypoint above sea level, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_SurfaceAltitude

```
<doc>
<summary>
The altitude of the waypoint above the surface of the body or sea level,
whichever is closer, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_SurfaceAltitude

```
<doc>
<summary>
The altitude of the waypoint above the surface of the body or sea level,
whichever is closer, in meters.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_BedrockAltitude

```
<doc>
<summary>
The altitude of the waypoint above the surface of the body, in meters.
When over water, this is the altitude above the sea floor.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_set_BedrockAltitude

```
<doc>
<summary>
The altitude of the waypoint above the surface of the body, in meters.
When over water, this is the altitude above the sea floor.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
```
value		
```
#### Return 
```
```
### Waypoint_get_NearSurface

```
<doc>
<summary><c>true</c> if the waypoint is near to the surface of a body.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_Grounded

```
<doc>
<summary><c>true</c> if the waypoint is attached to the ground.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_Index

```
<doc>
<summary>
The integer index of this waypoint within its cluster of sibling waypoints.
In other words, when you have a cluster of waypoints called "Somewhere Alpha",
"Somewhere Beta" and "Somewhere Gamma", the alpha site has index 0, the beta
site has index 1 and the gamma site has index 2.
When <see cref="M:SpaceCenter.Waypoint.Clustered" /> is <c>false</c>, this is zero.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_Clustered

```
<doc>
<summary><c>true</c> if this waypoint is part of a set of clustered waypoints with greek letter
names appended (Alpha, Beta, Gamma, etc).
If <c>true</c>, there is a one-to-one correspondence with the greek letter name and
the <see cref="M:SpaceCenter.Waypoint.Index" />.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_HasContract

```
<doc>
<summary>
Whether the waypoint belongs to a contract.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
```
### Waypoint_get_Contract

```
<doc>
<summary>
The associated contract.
</summary>
</doc>
```
#### Parameters 
```
this		Waypoint
```
#### Return 
```
Contract```
### WaypointManager_AddWaypoint

```
<doc>
<summary>
Creates a waypoint at the given position at ground level, and returns a
<see cref="T:SpaceCenter.Waypoint" /> object that can be used to modify it.
</summary>
<param name="latitude">Latitude of the waypoint.</param>
<param name="longitude">Longitude of the waypoint.</param>
<param name="body">Celestial body the waypoint is attached to.</param>
<param name="name">Name of the waypoint.</param>
<returns></returns>
</doc>
```
#### Parameters 
```
this		WaypointManager
```
```
latitude		
```
```
longitude		
```
```
body		CelestialBody
```
```
name		
```
#### Return 
```
Waypoint```
### WaypointManager_AddWaypointAtAltitude

```
<doc>
<summary>
Creates a waypoint at the given position and altitude, and returns a
<see cref="T:SpaceCenter.Waypoint" /> object that can be used to modify it.
</summary>
<param name="latitude">Latitude of the waypoint.</param>
<param name="longitude">Longitude of the waypoint.</param>
<param name="altitude">Altitude (above sea level) of the waypoint.</param>
<param name="body">Celestial body the waypoint is attached to.</param>
<param name="name">Name of the waypoint.</param>
<returns></returns>
</doc>
```
#### Parameters 
```
this		WaypointManager
```
```
latitude		
```
```
longitude		
```
```
altitude		
```
```
body		CelestialBody
```
```
name		
```
#### Return 
```
Waypoint```
### WaypointManager_get_Waypoints

```
<doc>
<summary>
A list of all existing waypoints.
</summary>
</doc>
```
#### Parameters 
```
this		WaypointManager
```
#### Return 
```
```
### WaypointManager_get_Icons

```
<doc>
<summary>
Returns all available icons (from "GameData/Squad/Contracts/Icons/").
</summary>
</doc>
```
#### Parameters 
```
this		WaypointManager
```
#### Return 
```
```
### WaypointManager_get_Colors

```
<doc>
<summary>
An example map of known color - seed pairs.
Any other integers may be used as seed.
</summary>
</doc>
```
#### Parameters 
```
this		WaypointManager
```
#### Return 
```
```
## KRPC
#### Classes
```
Expression
<doc>
<summary>
A server side expression.
</summary>
</doc>
```
```
Type
<doc>
<summary>
A server side expression.
</summary>
</doc>
```
### GetClientID

```
<doc>
<summary>
Returns the identifier for the current client.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### GetClientName

```
<doc>
<summary>
Returns the name of the current client.
This is an empty string if the client has no name.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### GetStatus

```
<doc>
<summary>
Returns some information about the server, such as the version.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### GetServices

```
<doc>
<summary>
Returns information on all services, procedures, classes, properties etc. provided by the server.
Can be used by client libraries to automatically create functionality such as stubs.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### AddStream

```
<doc>
<summary>
Add a streaming request and return its identifier.
</summary>
</doc>
```
#### Parameters 
```
call		
```
```
start		
```
#### Return 
```
```
### StartStream

```
<doc>
<summary>
Start a previously added streaming request.
</summary>
</doc>
```
#### Parameters 
```
id		
```
#### Return 
```
```
### SetStreamRate

```
<doc>
<summary>
Set the update rate for a stream in Hz.
</summary>
</doc>
```
#### Parameters 
```
id		
```
```
rate		
```
#### Return 
```
```
### RemoveStream

```
<doc>
<summary>
Remove a streaming request.
</summary>
</doc>
```
#### Parameters 
```
id		
```
#### Return 
```
```
### AddEvent

```
<doc>
<summary>
Create an event from a server side expression.
</summary>
</doc>
```
#### Parameters 
```
expression		Expression
```
#### Return 
```
```
### get_Clients

```
<doc>
<summary>
A list of RPC clients that are currently connected to the server.
Each entry in the list is a clients identifier, name and address.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### get_CurrentGameScene

```
<doc>
<summary>
Get the current game scene.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
GameScene```
### get_Paused

```
<doc>
<summary>
Whether the game is paused.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### set_Paused

```
<doc>
<summary>
Whether the game is paused.
</summary>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
```
### Expression_static_ConstantDouble

```
<doc>
<summary>
A constant value of double precision floating point type.
</summary>
<param name="value"></param>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
Expression```
### Expression_static_ConstantFloat

```
<doc>
<summary>
A constant value of single precision floating point type.
</summary>
<param name="value"></param>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
Expression```
### Expression_static_ConstantInt

```
<doc>
<summary>
A constant value of integer type.
</summary>
<param name="value"></param>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
Expression```
### Expression_static_ConstantBool

```
<doc>
<summary>
A constant value of boolean type.
</summary>
<param name="value"></param>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
Expression```
### Expression_static_ConstantString

```
<doc>
<summary>
A constant value of string type.
</summary>
<param name="value"></param>
</doc>
```
#### Parameters 
```
value		
```
#### Return 
```
Expression```
### Expression_static_Call

```
<doc>
<summary>
An RPC call.
</summary>
<param name="call"></param>
</doc>
```
#### Parameters 
```
call		
```
#### Return 
```
Expression```
### Expression_static_Equal

```
<doc>
<summary>
Equality comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_NotEqual

```
<doc>
<summary>
Inequality comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_GreaterThan

```
<doc>
<summary>
Greater than numerical comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_GreaterThanOrEqual

```
<doc>
<summary>
Greater than or equal numerical comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_LessThan

```
<doc>
<summary>
Less than numerical comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_LessThanOrEqual

```
<doc>
<summary>
Less than or equal numerical comparison.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_And

```
<doc>
<summary>
Boolean and operator.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Or

```
<doc>
<summary>
Boolean or operator.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_ExclusiveOr

```
<doc>
<summary>
Boolean exclusive-or operator.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Not

```
<doc>
<summary>
Boolean negation operator.
</summary>
<param name="arg"></param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Add

```
<doc>
<summary>
Numerical addition.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Subtract

```
<doc>
<summary>
Numerical subtraction.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Multiply

```
<doc>
<summary>
Numerical multiplication.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Divide

```
<doc>
<summary>
Numerical division.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Modulo

```
<doc>
<summary>
Numerical modulo operator.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
<returns>The remainder of arg0 divided by arg1</returns>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Power

```
<doc>
<summary>
Numerical power operator.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
<returns>arg0 raised to the power of arg1, with type of arg0</returns>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_LeftShift

```
<doc>
<summary>
Bitwise left shift.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_RightShift

```
<doc>
<summary>
Bitwise right shift.
</summary>
<param name="arg0"></param>
<param name="arg1"></param>
</doc>
```
#### Parameters 
```
arg0		Expression
```
```
arg1		Expression
```
#### Return 
```
Expression```
### Expression_static_Cast

```
<doc>
<summary>
Perform a cast to the given type.
</summary>
<param name="arg"></param>
<param name="type">Type to cast the argument to.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
type		Type
```
#### Return 
```
Expression```
### Expression_static_Parameter

```
<doc>
<summary>
A named parameter of type double.
</summary>
<returns>A named parameter.</returns>
<param name="name">The name of the parameter.</param>
<param name="type">The type of the parameter.</param>
</doc>
```
#### Parameters 
```
name		
```
```
type		Type
```
#### Return 
```
Expression```
### Expression_static_Function

```
<doc>
<summary>
A function.
</summary>
<returns>A function.</returns>
<param name="parameters">The parameters of the function.</param>
<param name="body">The body of the function.</param>
</doc>
```
#### Parameters 
```
parameters		
```
```
body		Expression
```
#### Return 
```
Expression```
### Expression_static_Invoke

```
<doc>
<summary>
A function call.
</summary>
<returns>A function call.</returns>
<param name="function">The function to call.</param>
<param name="args">The arguments to call the function with.</param>
</doc>
```
#### Parameters 
```
function		Expression
```
```
args		
```
#### Return 
```
Expression```
### Expression_static_CreateTuple

```
<doc>
<summary>
Construct a tuple.
</summary>
<returns>The tuple.</returns>
<param name="elements">The elements.</param>
</doc>
```
#### Parameters 
```
elements		
```
#### Return 
```
Expression```
### Expression_static_CreateList

```
<doc>
<summary>
Construct a list.
</summary>
<returns>The list.</returns>
<param name="values">The value. Should all be of the same type.</param>
</doc>
```
#### Parameters 
```
values		
```
#### Return 
```
Expression```
### Expression_static_CreateSet

```
<doc>
<summary>
Construct a set.
</summary>
<returns>The set.</returns>
<param name="values">The values. Should all be of the same type.</param>
</doc>
```
#### Parameters 
```
values		
```
#### Return 
```
Expression```
### Expression_static_CreateDictionary

```
<doc>
<summary>
Construct a dictionary, from a list of corresponding keys and values.
</summary>
<returns>The dictionary.</returns>
<param name="keys">The keys. Should all be of the same type.</param>
<param name="values">The values. Should all be of the same type.</param>
</doc>
```
#### Parameters 
```
keys		
```
```
values		
```
#### Return 
```
Expression```
### Expression_static_ToList

```
<doc>
<summary>
Convert a collection to a list.
</summary>
<returns>The collection as a list.</returns>
<param name="arg">The collection.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_ToSet

```
<doc>
<summary>
Convert a collection to a set.
</summary>
<returns>The collection as a set.</returns>
<param name="arg">The collection.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Get

```
<doc>
<summary>
Access an element in a tuple, list or dictionary.
</summary>
<returns>The element.</returns>
<param name="arg">The tuple, list or dictionary.</param>
<param name="index">The index of the element to access.
A zero indexed integer for a tuple or list, or a key for a dictionary.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
index		Expression
```
#### Return 
```
Expression```
### Expression_static_Count

```
<doc>
<summary>
Number of elements in a collection.
</summary>
<returns>The number of elements in the collection.</returns>
<param name="arg">The list, set or dictionary.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Sum

```
<doc>
<summary>
Sum all elements of a collection.
</summary>
<returns>The sum of the elements in the collection.</returns>
<param name="arg">The list or set.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Max

```
<doc>
<summary>
Maximum of all elements in a collection.
</summary>
<returns>The maximum elements in the collection.</returns>
<param name="arg">The list or set.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Min

```
<doc>
<summary>
Minimum of all elements in a collection.
</summary>
<returns>The minimum elements in the collection.</returns>
<param name="arg">The list or set.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Average

```
<doc>
<summary>
Minimum of all elements in a collection.
</summary>
<returns>The minimum elements in the collection.</returns>
<param name="arg">The list or set.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
#### Return 
```
Expression```
### Expression_static_Select

```
<doc>
<summary>
Run a function on every element in the collection.
</summary>
<returns>The modified collection.</returns>
<param name="arg">The list or set.</param>
<param name="func">The function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
func		Expression
```
#### Return 
```
Expression```
### Expression_static_Where

```
<doc>
<summary>
Run a function on every element in the collection.
</summary>
<returns>The modified collection.</returns>
<param name="arg">The list or set.</param>
<param name="func">The function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
func		Expression
```
#### Return 
```
Expression```
### Expression_static_Contains

```
<doc>
<summary>
Determine if a collection contains a value.
</summary>
<returns>Whether the collection contains a value.</returns>
<param name="arg">The collection.</param>
<param name="value">The value to look for.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
value		Expression
```
#### Return 
```
Expression```
### Expression_static_Aggregate

```
<doc>
<summary>
Applies an accumulator function over a sequence.
</summary>
<returns>The accumulated value.</returns>
<param name="arg">The collection.</param>
<param name="func">The accumulator function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
func		Expression
```
#### Return 
```
Expression```
### Expression_static_AggregateWithSeed

```
<doc>
<summary>
Applies an accumulator function over a sequence, with a given seed.
</summary>
<returns>The accumulated value.</returns>
<param name="arg">The collection.</param>
<param name="seed">The seed value.</param>
<param name="func">The accumulator function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
seed		Expression
```
```
func		Expression
```
#### Return 
```
Expression```
### Expression_static_Concat

```
<doc>
<summary>
Concatenate two sequences.
</summary>
<returns>The first sequence followed by the second sequence.</returns>
<param name="arg1">The first sequence.</param>
<param name="arg2">The second sequence.</param>
</doc>
```
#### Parameters 
```
arg1		Expression
```
```
arg2		Expression
```
#### Return 
```
Expression```
### Expression_static_OrderBy

```
<doc>
<summary>
Order a collection using a key function.
</summary>
<returns>The ordered collection.</returns>
<param name="arg">The collection to order.</param>
<param name="key">A function that takes a value from the collection and generates a key to sort on.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
key		Expression
```
#### Return 
```
Expression```
### Expression_static_All

```
<doc>
<summary>
Determine whether all items in a collection satisfy a boolean predicate.
</summary>
<returns>Whether all items satisfy the predicate.</returns>
<param name="arg">The collection.</param>
<param name="predicate">The predicate function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
predicate		Expression
```
#### Return 
```
Expression```
### Expression_static_Any

```
<doc>
<summary>
Determine whether any item in a collection satisfies a boolean predicate.
</summary>
<returns>Whether any item satisfies the predicate.</returns>
<param name="arg">The collection.</param>
<param name="predicate">The predicate function.</param>
</doc>
```
#### Parameters 
```
arg		Expression
```
```
predicate		Expression
```
#### Return 
```
Expression```
### Type_static_Double

```
<doc>
<summary>
Double type.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Type```
### Type_static_Float

```
<doc>
<summary>
Float type.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Type```
### Type_static_Int

```
<doc>
<summary>
Int type.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Type```
### Type_static_Bool

```
<doc>
<summary>
Bool type.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Type```
### Type_static_String

```
<doc>
<summary>
String type.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
Type```
## Drawing
#### Classes
```
Line
<doc>
<summary>
A line. Created using <see cref="M:Drawing.AddLine" />.
</summary>
</doc>
```
```
Polygon
<doc>
<summary>
A polygon. Created using <see cref="M:Drawing.AddPolygon" />.
</summary>
</doc>
```
```
Text
<doc>
<summary>
Text. Created using <see cref="M:Drawing.AddText" />.
</summary>
</doc>
```
### AddLine

```
<doc>
<summary>
Draw a line in the scene.
</summary>
<param name="start">Position of the start of the line.</param>
<param name="end">Position of the end of the line.</param>
<param name="referenceFrame">Reference frame that the positions are in.</param>
<param name="visible">Whether the line is visible.</param>
</doc>
```
#### Parameters 
```
start		
```
```
end		
```
```
referenceFrame		ReferenceFrame
```
```
visible		
```
#### Return 
```
Line```
### AddDirection

```
<doc>
<summary>
Draw a direction vector in the scene, from the center of mass of the active vessel.
</summary>
<param name="direction">Direction to draw the line in.</param>
<param name="referenceFrame">Reference frame that the direction is in.</param>
<param name="length">The length of the line.</param>
<param name="visible">Whether the line is visible.</param>
</doc>
```
#### Parameters 
```
direction		
```
```
referenceFrame		ReferenceFrame
```
```
length		
```
```
visible		
```
#### Return 
```
Line```
### AddPolygon

```
<doc>
<summary>
Draw a polygon in the scene, defined by a list of vertices.
</summary>
<param name="vertices">Vertices of the polygon.</param>
<param name="referenceFrame">Reference frame that the vertices are in.</param>
<param name="visible">Whether the polygon is visible.</param>
</doc>
```
#### Parameters 
```
vertices		
```
```
referenceFrame		ReferenceFrame
```
```
visible		
```
#### Return 
```
Polygon```
### AddText

```
<doc>
<summary>
Draw text in the scene.
</summary>
<param name="text">The string to draw.</param>
<param name="referenceFrame">Reference frame that the text position is in.</param>
<param name="position">Position of the text.</param>
<param name="rotation">Rotation of the text, as a quaternion.</param>
<param name="visible">Whether the text is visible.</param>
</doc>
```
#### Parameters 
```
text		
```
```
referenceFrame		ReferenceFrame
```
```
position		
```
```
rotation		
```
```
visible		
```
#### Return 
```
Text```
### Clear

```
<doc>
<summary>
Remove all objects being drawn.
</summary>
<param name="clientOnly">If true, only remove objects created by the calling client.</param>
</doc>
```
#### Parameters 
```
clientOnly		
```
#### Return 
```
```
### Line_Remove

```
<doc>
<summary>
Remove the object.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_get_Start

```
<doc>
<summary>
Start position of the line.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_Start

```
<doc>
<summary>
Start position of the line.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Line_get_End

```
<doc>
<summary>
End position of the line.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_End

```
<doc>
<summary>
End position of the line.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Line_get_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Line_get_Thickness

```
<doc>
<summary>
Set the thickness
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_Thickness

```
<doc>
<summary>
Set the thickness
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Line_get_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
ReferenceFrame```
### Line_set_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		ReferenceFrame
```
#### Return 
```
```
### Line_get_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Line_get_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
#### Return 
```
```
### Line_set_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Line
```
```
value		
```
#### Return 
```
```
### Polygon_Remove

```
<doc>
<summary>
Remove the object.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_get_Vertices

```
<doc>
<summary>
Vertices for the polygon.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_set_Vertices

```
<doc>
<summary>
Vertices for the polygon.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		
```
#### Return 
```
```
### Polygon_get_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_set_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		
```
#### Return 
```
```
### Polygon_get_Thickness

```
<doc>
<summary>
Set the thickness
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_set_Thickness

```
<doc>
<summary>
Set the thickness
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		
```
#### Return 
```
```
### Polygon_get_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
ReferenceFrame```
### Polygon_set_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		ReferenceFrame
```
#### Return 
```
```
### Polygon_get_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_set_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		
```
#### Return 
```
```
### Polygon_get_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
#### Return 
```
```
### Polygon_set_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Polygon
```
```
value		
```
#### Return 
```
```
### Text_static_AvailableFonts

```
<doc>
<summary>
A list of all available fonts.
</summary>
</doc>
```
#### Parameters 
#### Return 
```
```
### Text_Remove

```
<doc>
<summary>
Remove the object.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_get_Position

```
<doc>
<summary>
Position of the text.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Position

```
<doc>
<summary>
Position of the text.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Rotation

```
<doc>
<summary>
Rotation of the text as a quaternion.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Rotation

```
<doc>
<summary>
Rotation of the text as a quaternion.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Content

```
<doc>
<summary>
The text string
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Content

```
<doc>
<summary>
The text string
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Font

```
<doc>
<summary>
Name of the font
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Font

```
<doc>
<summary>
Name of the font
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Size

```
<doc>
<summary>
Font size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Size

```
<doc>
<summary>
Font size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_CharacterSize

```
<doc>
<summary>
Character size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_CharacterSize

```
<doc>
<summary>
Character size.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Style

```
<doc>
<summary>
Font style.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
FontStyle```
### Text_set_Style

```
<doc>
<summary>
Font style.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		FontStyle
```
#### Return 
```
```
### Text_get_Alignment

```
<doc>
<summary>
Alignment.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
TextAlignment```
### Text_set_Alignment

```
<doc>
<summary>
Alignment.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		TextAlignment
```
#### Return 
```
```
### Text_get_LineSpacing

```
<doc>
<summary>
Line spacing.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_LineSpacing

```
<doc>
<summary>
Line spacing.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Anchor

```
<doc>
<summary>
Anchor.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
TextAnchor```
### Text_set_Anchor

```
<doc>
<summary>
Anchor.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		TextAnchor
```
#### Return 
```
```
### Text_get_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Color

```
<doc>
<summary>
Set the color
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
ReferenceFrame```
### Text_set_ReferenceFrame

```
<doc>
<summary>
Reference frame for the positions of the object.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		ReferenceFrame
```
#### Return 
```
```
### Text_get_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Visible

```
<doc>
<summary>
Whether the object is visible.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
### Text_get_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
#### Return 
```
```
### Text_set_Material

```
<doc>
<summary>
Material used to render the object.
Creates the material from a shader with the given name.
</summary>
</doc>
```
#### Parameters 
```
this		Text
```
```
value		
```
#### Return 
```
```
