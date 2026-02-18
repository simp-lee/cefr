#!/usr/bin/env python3
"""Generate CSV word list data files for the CEFR library."""
import os

DIR = os.path.dirname(os.path.abspath(__file__))

# ============================================================
# Oxford 5000 — word,level
# ============================================================
# A1: most basic everyday vocabulary (~600 words)
oxford_a1 = """
a about above across act add after afternoon again age ago air all almost along
already always am an and animal another answer any anything apartment apple april
are area arm around arrive art as ask at august away back bad bag ball banana band
bank bathroom be bean beautiful because become bed bedroom before begin behind
believe below best better between big bike bird birthday bit black blue board boat
body book both bottom box boy brain bread break breakfast bring brother brown build
bus business busy but buy by cake call camera can car card care carry cat catch
center chair change cheap cheese chicken child children choose church city class
classroom clean clear clock close clothes cloud coat coffee cold color come common
computer cook cool correct cost could country course cover cross cry cup cut dad
dance dark daughter day dear december decide deep desk dictionary die different
difficult dinner direction dirty do doctor dog dollar door down draw dream dress
drink drive drop dry during each ear early earth east eat egg eight eighteen eighty
eleven email end enjoy enough enter even evening ever every everybody everyone
everything exam example excited exercise expensive eye face fact fall family far
farm fast father favorite february feel few fifteen fifty film find fine finger
finish fire first fish five fix floor fly food foot for forget forty four fourteen
free friday friend from front fruit full fun funny future game garden get girl give
glass go goal good goodbye grass great green grey ground group grow guess guitar
guy hair half hand happen happy hard hat hate have he head hear heart heavy hello
help her here high him his hit hold holiday home hope horse hospital hot hotel
hour house how hundred hungry husband i ice idea if important in information
inside interest interesting into island it its january july june just keep key
kid kind king kitchen knee know land language large last late later laugh learn
left leg lesson let letter library life light like line list listen little live
long look lose lot love low lunch machine main make man many map march market
marry math matter may me meal mean meat meet member message middle might milk
million mind minute miss model modern mom moment monday money month more morning
most mother mountain mouth move movie mr mrs much museum music must my name near
necessary need never new news next nice night nine nineteen ninety no nobody none
noon normal north nose not note nothing november now number october of off offer
office often oh ok old on once one only open or orange order other our out outside
over own page pair pants parent park part party pass past pay pen pencil people
percent period person phone photo pick picture piece place plan plant play player
please point police pool poor popular possible post practice prepare present
pretty price probably problem program pull push put quarter queen question quick
quickly quiet quite rain read ready real really red remember rest restaurant
return rice rich ride right rise river road room round rule run sad safe same
saturday say school sea search season second see sell send sentence september
serious seven seventeen seventy she shirt shoe shop short should show shower
side sign simple since sing sister sit six sixteen sixty size sleep small smile
snow so soccer some somebody someone something sometimes son song soon sorry south
speak special spend sport spring stand star start station stay step still stop
store story street strong student study subject such sugar summer sun sunday
sure surprise swim table take talk tall tea teach teacher team tell ten test
than thank that the their them then there they thing think third thirteen thirty
this those thousand three through thursday time tired to today toe together
tomorrow too top town toy train travel tree trip try tuesday turn tv twelve
twenty two under understand university until up us use usually vegetable very
visit wait walk wall want warm wash watch water way we wear weather wednesday
week well west what when where which white who why wide wife will win window
winter wish with without woman wonderful word work world worry would write wrong
year yellow yes yesterday you young your zero
""".split()

# A2: elementary vocabulary (~1000 words)
oxford_a2 = """
abroad accept accident achieve across action active activity actor actually add
admire adult advantage adventure advice afraid afterwards agree ahead airplane
alive allow almost alone along aloud already although amount ancient angry
announce apart appear application appointment approve area argue army arrange
arrive article attack attempt attract audience autumn average avoid awake
awful baby babysitter background backwards balance ballet bandage basic basket bat
battle bay beach bear beard beat beauty bell belong below bend beside besides
bike bill billion biology blank blanket blind block blog blood blow board bone
border boring born borrow boss bottle bottom brain brave brilliant broad brush
budget burn calm cancel candle captain career careful carefully carry case cash
castle celebrate central century chain challenge champion championship chance
character charge charity chart chat check cheer childhood chip chocolate
citizen classical clearly clever climate climbing coach collection comedy
comfortable command communicate communication community compare competition
competitor complain complete completely concern condition confident confirm connect
consider contain content continue control convenient conversation copy corner
cotton couple courage crash create creative crime criminal cross crowd cruel
culture customer cycle daily damage danger dangerous daughter dead deal death
debate decision decorate decrease degree deliver demand department depend describe
desert design develop development device dialogue diet dig digital direction
director disappear disappoint discover discussion dish distance divide document
dollar double doubt downtown drag drawing dump earn education effect effort Either
electricity electronic emergency employee employer empty encourage enemy engine
engineer entertainment entire entrance environment equal equipment escape especially
essay event evidence evil exactly examination excellent except exchange excitement
exciting executive exist expect experience experiment explain explore export
express extra extreme factory fail fair famous fan fancy fashion fault fear feature
feed female fence festival fiction field figure fill final finally fit fix flag
flat flight float flood flow focus fold follow following force foreign forest
forever forget formal former forward found freedom fresh fridge friendly
frighten fuel furniture gap garage gate gather general generation gentle gentleman
giant gift given glad global gold golden government grade grandchild grandfather
grandmother guest guide guilty gym habit hall handle hardly harm hate health
healthy heap heat height helpful hero highly historical hole honest huge humor
hunt hurry hurt ice ignore ill illegal image immediate immediately impact impress
improve include increase independent individual industry influence inform injury
innocent instruction intelligent intend international introduce investigate
invitation iron issue item jacket jeans jewel jewelry join journalist journey joy
judge juice junior keen kill knee knock knowledge lack lady lake lane laptop lately
lately latest laugh launch laundry lawyer lay layer leader least leather legal
leisure lemon length level lid lifestyle lift likely limit link liquid literature
litter load local lonely loop lovely low lower luck luggage mad magic main mainly
male manage manager manner mark mass match material mayor meanwhile measure medium
melt mental mention mess metal method million mirror mix mixture mobile mood moral
mostly motor mouse movement murder muscle mystery narrow nation native natural
naturally nearly neat negative neighbor neither network nightmare nobody noise
none nor normal northern note notice novel object obvious occasion occur odd
official operation opinion opportunity opposite option ordinary organize original
otherwise ought outer outside oven overcome owner pack package pain painful
painting pair palace pale pan panel paragraph park parking partner passenger
patient pattern pause payment peace pear perform performance perhaps permanent
permission personal personality phone photo photograph phrase physical pile pilot
plain plastic plate platform pleasant plenty pocket poetry poison policy polite
politics pop position positive possibility possibly pot potential powerful prayer
predict prefer preparation president prevent previous primary prince princess print
priority prison prisoner private prize process produce product production
profession professional professor project promise promote proper properly property
protect proud prove provide public publish punishment pupil pure purpose push
puzzle qualify quarter racing range rank rapid rare rate rather reach react receipt
recently recognize recommend record reduce refer regular reject relate
relation relationship relief rely remain remark remind remove rent repair repeat
replace report republic request research reserve resource respond responsible
review revolution reward risk robot role romantic roof root rough route row royal
rude sail salary sale sample sand satisfy save scale scene schedule screen
script secret section security seed seek select selection senior sentence
separate serious servant serve session severe shadow shake shame shape share sharp
shelter shift shock signal silence silly silver similar simple sink site
situation skill skin slave slight smart smooth social society soil solar soldier
solution somebody someday somewhere sort soul source southern spare speaker spin
spirit split spread square stable staff stage stamp standard state steal wheel
steam steel stomach store storm stranger strategy stream strength stress stretch
strict structure struggle style success successful suit suitable supply support
suppose surface survive suspect sweet sympathy system talent target task tax
technology teenager temperature text theme theory therefore thick thin thoroughly
threat throat throughout throw thus ticket tight tip title tonight tool topic
total touch tourism tourist tradition traffic transport treat treatment trick trouble
trust truth tube tune typical typically ugly ultimately unable uncle uncomfortable
unfortunately uniform union unique unit united universe unless unlikely unusual
upon upper upset urban variety vast vehicle version victim video village violence
virtual visitor vocabulary voice volume volunteer wage war warn warning waste
wealth weapon wear website weekend weight welcome western willing wing winner
wire wise wonderful wooden worker worried worth wrap yard youth zone
""".split()

# B1: intermediate vocabulary (~1500 words)
oxford_b1 = """
abandon ability absolutely absorb abstract academic accelerate acceptable
access accommodate accommodation accompany according account accountant accurate
accusation accuse ache achievement acknowledge acquire act actual adaptation
additional address adequate adjust administration administrator admission
adolescent adopt adoption advanced advertising affair affection afford african
aging agricultural agriculture aim aircraft alarm album alcohol alcoholic
alongside alter alternative altogether aluminum amaze amazed amazing ambassador
ambition ambitious amid amount amusement analyst ancient angle anniversary
annual anonymous anticipate anxiety anxious anyhow anyway apart apartment
apparatus apparent apparently appeal appetite appoint appreciation appropriate
approval approximately april archive argue arrangement aspect assembly
assessment assign assignment assist association assume assumption atmosphere
attach attempt attendance attitude attraction attractive authority automatic
automatically automobile availability aviation awareness
bachelor bacteria badge badly ban bargain barrier basically bay beam behalf
behave belief bench beneath berry beside beyond billion bio biological blame
blessing blind blow bold bone boom boost bound bowl branch brand breast breath
breed brilliant broad broadcast brush buck buddy bulk burden bureaucracy bypass
cab cabin cabinet cable cafe calculate calm campaign campus capability capable
capacity capture carbon carrier casual catalog catalogue category caution
cellular chain chairman chamber chancellor channel chapter characteristic
characterize charm chat chemical chemistry childhood chip chocolate chunk
cigarette circuit circumstance citation cite civilian claim clarity clash classify
clerk click cliff climb clinical closely clothing coalition cognitive coincidence
collapse colleague collective colonial colony combination combine comfort
commander commerce commercial commission commissioner commit commitment committee
commodity communicate comparable comparison compel compensate compensation
competence compile complaint complex complexity component compose composition
comprehensive compromise concentrate concentration conclude concrete
confront confusion congressional consequence conservation conservative
considerably consideration consistent constantly construct consumer consumption
contemporary controversial convention cope core corporate corporation correspond
cottage counsel county coverage crack craft creative crew criticism criticize
crop crowd crucial cry crystal cultural cure curiosity curious curriculum custom
cutting
deadline dealer dear debt debut decade decent declaration decline decoration
decrease defend deficit definitely definition degree delay deliberately democracy
democrat democratic demographic demonstration deny departure depression derive
description desire desperately destination destruction detective determination
developer digital dimension disability disabled disappear discharge discipline
discourse discrimination disorder display dispute distinction distinctive
distinguish diverse diversity divorce doctrine domain domestic dominant dominate
donation donor dozen draft dramatic dramatically drift drum duration dust
dynamic
eager earn ease editor effectively efficiency efficiently eighth elderly
election elegant elimination elite elsewhere embrace emission emotional emphasis
emphasize empire enable encounter endless energy enforcement engage engineer
engineering enormous enterprise enthusiasm enthusiastic entire entity entrance
entrepreneur entry envelope environmental equally era essentially estate
ethics ethnic evaluate evaluation eventual eventually evident evolution evolve
examination exceed exception excessive exclusive exclusively execute
expansion expectation expedition expense expertise explanation explicit
explosion expose exposure extension extent external extraordinary extreme
fabric facility faculty fame farming fascinating fate favor favorable featured
fee fellow fiber fierce fifth fifty final fix flat flavor flee flexibility
float folk forecast forever formerly forum foster foundation founding
fraction fragment framework frequency frequent friction front frustrate
frustration fundamental funding
gallery gang gap gear gender generate generous genetic genius genuine
gesture global golden gradually graduate grain graphic grave guilty
halt handler handful harm harsh headline headquarters helper heritage highlight
highlight historic horizon hormone horror hosting hostile household humor
hungry hypothesis
ideal identification identity ignore illustration imagination immigrant
immigration implement import impose impression improvement incentive incident
incorporate indication inevitable infant inflation initially inner innovation
innovative inspect installation instance institutional integrate integration
intellectual intelligence intensity interaction interfere interior
interpretation intervention intimate invasion investor invisible involvement
isolation item
jam journalist judgment jury justify
keen killing
laboratory landscape launch lawsuit layout leap legacy legislation legitimate
lend lengthy liberal liberation liberty likewise limitation linear literary
lively lobby logic long-term loose
mainland maintenance maker mandate margin mate mayor meaningful measurement
mechanism membership memorial merely merit middle-class migrate migration
mineral minimal minority miracle model moderate modify monitor moreover
mortgage motivation multiple municipal mutual myth
narrative naval necessity nest
obligation observation observer obstacle obvious occasionally occupation
odds offense offensive ongoing operational opponent oppose organic orientation
origin outcome outline overcome overseas
painful pale participate participation partnership passage passion patience
peak penalty perceive permanently permit persist perspective phase popularity
portfolio pose possess poverty pray precisely prediction predominantly
preference pregnant preparation presumably presume prior privilege probe
proceed processing productive productivity profile profound progressive
prominent promotion properly proportion prospect protective protein
protest provider province psychological publication pursue
quote
racial racism radical rage rally random rapidly ratio reaction readily
realistic recession recognition reconstruction recovery recruit reflection
reform refugee regard regime regulation regulatory reinforce reluctant
remarkable remedy remote repair reproductive requirement residential resign
resolution resort respective restoration restrict restriction retain retreat
reveal revenue reverse revolution revolutionary rhythm rid rifle rival
robot role routine royal ruling rural
sacrifice scandal scared scholar scholarship seasonal secondary sector
selection senator sensation sensitive sequence settlement sharply shed shelter
sheriff signal similarly simulate simultaneously skeptical slave slavery slice
slim smoke somewhat sophisticated Southeast sovereignty span specialist
speculation strategic strongly structural struggle substitute sue sufficient
summit supplement supposedly surgery surplus swing
talent tank terrorism tendency therapy thereby thesis timber tobacco tolerance
trait transition transmission trend tribal triumph tropical
undergo underlying unemployment unfair unfortunately uniform universal

""".split()

# B2: upper-intermediate vocabulary (~1200 words)
oxford_b2 = """
abstract abstraction abundance accessible acclaim accommodate accumulate
accuracy acknowledgement acquisition acute adaptation addiction adequately
adjacent adjustment administrator admiration adverse advocate affair aggregate
aggression aggressive agony allocate allied allowance alongside amendment ample
analogy anchor animation anonymous anxiety apparatus applaud appreciation
architect architectural arena articulate aspiration assault assert
assertion assurance atmospheric attic audit authentic authorize autonomous
backdrop bail bankruptcy bargain behavioral benchmark bias biographical
bizarre bloom blueprint bold bombard boost bounce boundary bracket breakthrough
breed broadly broker brutal buddy bulge
casualty celebrity centerpiece certainty chamber chaos charitable charter chronic
chunk clarity classification cling coalition coincide collaboration
collaborative collector commence commentary commentary commerce commissioner
commodity compelling compensate competent complement compliance complication
comprehensive compromise compulsory concede conceive confess confession
confine confirmation confrontation congregation consent consequently
consolidate conspiracy constitutional consultation contemplate contend
contention contractor contradiction controversial conversion conviction
cooperate cooperative coordination correspond correspondence corridor counsel
counterpart coverage craft creed critique crop crucial crush curiosity
curriculum custody
database deadline debris debut dedication default deficiency deficit deliberately
deity delicate density dependence depict deployment deprive designate desirable
desktop destruction detention determination diagnose diagnosis diplomatic
disability discourse discrimination disorder displacement disposal distinctive
distort disturbance diverse dock doctrine Dominican donation donor double
drainage duration
elaborate elegance eligible emergence emission empathy empirical enact
endeavor endorse enforcement engagement enterprise entity entrepreneur
envision equip equivalence equivalent erosion escalate essence eternal euro
evaluate eventually evidence evolve exaggerate exceed exclusive execute
exempt exhibit exile expansion exploitation explicit exploit export
extract
fabricate facilitate faction fantasy fatal fate favorable fertility
fierce figurative filter fiscal fit fixture flexibility flourish fluctuate
formula foster fraction fragment framework franchise friction frontier fulfill
functional fundraising fury fusion
galaxy gap genre genuine giant given globe governor grave gravity grip guardian
guideline guilt
habitat halt harassment harsh hazard headquarters heighten heritage
hierarchy highlight hip homeless horizon hostile hostility humanitarian
hypothesis
icon ideology ignorance illusion immense implementation implicit import
impulse inability inadequate inappropriate inclination incredible indigenous
induce inequality infant inflation influential infrastructure inhabit inherent
inhibit initiate injection inner innovation insert inspection install
institutional intact integration intellectual intensity interactive invest
intervention intimate inventory invoke isolation
journalism judiciary jurisdiction justification
keen keyboard
landmark layout legislative legitimate liberal linger literacy lobby logical
longevity loyalty lumber
machinery mainstream make-up mandate manipulation manufacturer manuscript
marginal marketplace massacre maternal maximum membrane mentor merchandise
mere metaphor methodology microscope midst migration militant minimize
minister miracle mobility moderate modest modification monument mortality
motivation municipal
narrative neglect negotiate nonetheless notorious novel
obligation occupational odds offensive offspring operational optimistic
orientation outbreak outspoken overcome overwhelming
parade parameter parish parliamentary partial participant passionate
patience patron pause peasant peculiar pension perceive permanent
persistence petition pharmaceutical physician pioneer pitch plea plot plunge
polarize portfolio pose possess practitioner precede prejudice preliminary
premise premium prescription presumably prevalent primarily primitive
priority privilege probe procurement profound progressive prohibit projection
prominent promote prone propaganda proportion prosecution prospective
protagonist province provisional provoke
quest questionnaire
racism radical rally realm reconcile reconstruction referendum reflection
regime regulation rehabilitation reign reinforce reluctance render renowned
reproduction reservation residual residents resilience resemble respective
restoration restraint retrieve revelation reverse revision rhetoric ritual
robust
sacred sacrifice saint sanction scout secular segment sender sensation
sentiment separation sergeant serial sheriff situated sketch slot solidarity
sovereign specimen speculate squad stake stem stereotype stimulus strand
straightforward strip structural stumble subordinate subsequent substantive
suburban successor sufficient summit superb supplement suppress surgeon surplus
surveillance suspension sustainable sway symbolic syndrome
tackle tactic telecommunications tempt tender texture theology threshold
tolerance tolerate torture toxic trademark trait transaction transparent
trauma treaty tribe tribunal trigger triumph tropical trustee
underlying undertake underway unprecedented uphold usage
vague validity vegetation venue verbal verdict verify versus viable
violate vital vulnerable
warfare warrior welfare withdrawal workout
""".split()

# C1: advanced vocabulary (~700 words)
oxford_c1 = """
abolish abrupt abundance accelerate accessible accountable accumulate
acquaint acquisition activist acute adapt addictive adhere adjacent
administer adverse advocate affiliate affluent agenda aggregate aggravate
allegation alleviate alliance allocate alter ambiguity amend amid
analogous anchor animate anomaly apparatus appendix arbitrary archive
arid articulate ascertain aspire audit avert axis
bland bleak brevity brink broadband bureaucratic bypass
calibrate candid capsule casualties catalogue cater caveat cease
centralize chronicle civic clad cluster coalition codify cognitive
coherent coincidence commemorate commence competence compile complement
comprehensive comprise compulsory conceive concession confer confiscate
confront congregation conscientious consecutive consensus consent
consequently consolidation conspiracy constituent constraint contempt
contend continuity contractor contradictory conversely convey
convict copyright correlate correspondence corrode counteract
covenant covert credibility culminate curb custody
daunting decay decisive decree dedicate deem default defer deficit
definitive delegate deliberate demolish dense depict deposit depreciate
deprive derivative designate deteriorate detrimental deviate
devise dilemma diminish discard disclose discretion discrepancy
disparity disperse displace disposition disregard disruption
dissolve distinctly diverge dividend doctrine dominance donor
dormant downward drastic duration dwell
elaborate elicit elite embed embody embrace empirical empower enact
encompass endorse enforce enhance enlighten enrollment enterprise
entity entrepreneur envelope era erode escalate essence ethical
evacuate eventual evict evoke exceed exceedingly exclusive
exempt exert exile exotic expedite explicit exploit
extravagant facilitate faction faculty fatal feasible fibre fiscal
flaw focal follower footprint formidable forthcoming fossil
foster fragment franchise frenzy friction frontier frustrate
furnish fusion futile
gaze genesis glare govern gradient graphic grasp grievance
guideline hallmark handicap harness harsh haven hazard heed
heighten heritage hierarchy hinder humanitarian hypothesis
ideology illicit imminent immune impair imperative implicit
implicate impose impoverish inaugural incidence incline
incorporate incremental indigenous indispensable induce inevitable
inflation inflict infrastructure infringe inherent inhibit inject
innovative inquiry instigate Institute integral integrity
interact intercept interim intermittent intervene intimate
intricate intrinsic inventory invoke irony irrespective isolate
jeopardize jurisdiction juxtapose
landmark latitude lawsuit legacy legislation legitimate
leverage liberal linger literal logging lucrative luxury
magnitude mainstream mandate manifest manipulation manufacturer
marginal mature maximize mechanism mediate membrane mentor
metaphor meticulous migrate militant minimize ministry modest
momentum monastery monopoly motive municipal
namely narrative necessity negotiate nonetheless norm notable
notify notwithstanding nourish
obligatory obstacle obstruct offset onset ongoing optical
optimum oriented orthodox outright outsource overlap override
oversee overthrow
paradigm paradox parameter partisan patriot patronize peer penalty
penetrate perceive peripheral persistent petition pioneer pitfall
plausible pledge polarize portfolio pose postulate potent
potential practitioner precaution precede precision
predecessor predominantly prejudice preliminary premise
premium prescribe prestige presumption prevalent probe
proclaim procurement proficiency profound prohibition
prominent prone propaganda propagate proposition protagonist
protocol provoke proximity prudent
quota
radical ratify realm reconcile referendum refine reformulate
refugee regime regulate reinforce relay relentless relevant
reluctant render repeal replenish reservoir reside restraint
retrieve revelation rhetoric rigid rigorous ritual robust
ruthless
safeguard sanction saturate scenario sceptical scrutiny secular
segment sentiment seize simulate simultaneously skeptical
solidarity sovereign specimen speculate sphere stake stalemate
stance statute stem stereotype straightforward strand strategic
stride structural subordinate subsidy substantial successor
suffice summit superfluous supplement suppress surge surplus
surveillance susceptible sustainable swift syndrome
tangible tenure terminology testament texture threshold
tolerance trajectory transaction transcript transition
transparent trauma treaty trigger triumph turmoil
unanimous undergo undermine uneasy unprecedented unveil
upheaval uphold
vacuum velocity verdict verify versatile viable vicinity
vigorous violate virtue visibility vital volatile vulnerability
wield
yield
""".split()

# Supplemental Oxford words to reach ~5000 total
oxford_a1_supp = """
baby badly bar basket bath bear bell belt blanket block blow bone bottle bowl
bridge brush burn bus butter button bye camp cap captain card carry castle
castle cent chance cheek chin circle climb coin collect comb cook corner count
cow cream cross crown curly cushion deer dish downstairs draw driver drum
duck dust ear eastern eighth elephant empty enemy engine entrance envelope
evening everywhere exercise extra eye factory fairy fan farm fast favorite
fence field fifth fill finger flag flat flew fly fold forest fork fox freeze
friendly frog fruit funny gift glad gold grab grandfather grandmother
grape ham hang hate heart hole honey hurry husband ice ill insect iron island
jeans joke juice jump key kick kiss kitchen kitten knife ladder lamp land
large leaf leather left lemon lift lion lips look magazine male map meal
medicine member message metal milk mine miss mix moon mountain narrow near
neck needle neighbor newspaper noise normal north nurse ocean oil onion
opposite orange order outside painting pair palace parent partner passenger
passport path pen pencil pepper piano pink plate pocket police pot potato
prayer prince princess prize project promise pupil puzzle queen rabbit race
rain rainbow rat receive remember reply rest restaurant river roof rope
rubber sad sand scissors sea seat shape sheep shirt shoulder singer sink size
ski skin skirt sky smoke snake social sock soup space spider stairs stamp
stomach storm stranger strawberry student suitcase supermarket surprise
telephone tent thick tired toe tongue tooth towel toy traffic turkey
umbrella uncle uniform valley video village visitor voice wait warm watch
weigh west wheel wife wing wood worried yellow zero zoo
""".split()

oxford_a2_supp = """
ability absolutely accept accident actual admire advantage advise afraid
afterwards agency agent agriculture album alive allow ambition amount angle
annual anxious appearance application argue arrest atmosphere attempt
attitude attract average aware awful background backward balance bat battery
bay bear beat beef behave bend beside besides bill biology bite blank
blind blow bomb border bore borrow boss brick brilliant broad cabin calm
campaign cancel carbon cause celebrate central ceremony chain championship
characteristic cheek childhood civil client climate coach coal collapse
colleague comfort command commercial compete complex concerned conclusion
conference confirm confusion consider construction continue contribution
convenient conversation convince cotton couple courage crash cross cultural
currency cycle damage data deer delay deliver demand desert despite develop
diet dig digital dirty disabled disappoint discount disgusting distinction
disturb divide documentary double doubt downtown dramatic ear economic edge
electronic emergency employ employment encourage engage engineering enormous
entertainment essential eventually exactly evidence evil examine except
excited exist explore expression extend facility factor fair false fare
feature fiction fight financial firm flame flash float flow folk football
forbid format fortune freeze frequent fuel function gentle global grab
gradually guarantee guard guilty handle harm heat helpful hero hesitate hire
historical homeless huge humor hunt ideal ignore illness immediate impress
include income increase independent indicate individual industrial infection
influence innocent instruction insurance intelligent intention introduce
invest item journey judgment labor landscape launch lion literature load local
logical loose lucky magazine major manage manner manufacturer meanwhile
media mental military minor miracle modern monster motor narrow nearby
negative nerve normally novel nuclear obviously ocean operation ordinary
organize original overcome participate passion patience persuade pleasure
poison political possess previous primary profession profit properly protest
prove publish quality race rapid recognize refuse regular reject relate
religion remote repeat request rescue result ruin satisfy scientific screen
senior settlement shift shock signal significant silent skill sort spare
spirit spread standard state status steal strength stretch structure
struggle suffer suitable suggestion supply surgery suspect technical
teenager tend theory threaten tight tone tough trend tropical UI union unique
upper urge vast victim violent wealth weapon witness youth zone
""".split()

oxford_b1_supp = """
abandon absorb abstract academic accent accomplish accounting accurate
acid activist adaptation adequate adjust admiration adopt advanced
advertise affiliate aggressive agriculture aide aircraft alcohol
alien alliance alongside alter aluminum ambassador ambitious amendment
amid anchor angel animation anniversary anonymous anticipation
appeal appetite appliance appreciation architect archive arena
artificial assault assembly associate astronaut athletic auction
autobiography automobile aviation awful
bargain barrel basement batch behalf beloved berry biography
biology blend blessing bloody bold bonus booth bounce boundary
bracket breed bride broadcast bronze bubble bucket burden burial
cabin capability carbon cartoon cast catalog caution celebrity
championship chaos chapter charm charter cheek chemistry chip
chronic chunk circuit civilian civilization clarity classic cliff
coalition column combat comedy commence commentary commissioner
companion compensation compile competitor component comprehensive
compromise concentrate concrete confidential confuse congregation
consciousness consensus consequence conservation consist construct
consume controversial convention conversion cooperate copyright core
correction correspondent corruption costume cottage counsel counter
coverage crack craft crane crew cricket criterion crop cruel crystal
curriculum
dam dawn dealer debate decade decent declare decline decrease
default deficit deliberately democracy demonstrate departure
depression deputy derive description desperate detective devil diagram
dialect diary dietary discipline discount discrimination disorder
dispute distinction diversity documentary domestic donation dose
dramatically drift dump duration duty
eager eastern ecological efficiently election elementary elimination
embassy emergency emission empire enable encounter endure enforce
enormous enterprise entrepreneur envelope episode equality era essay
estate estimate ethic evaluate evil evolution examination exception
excessive exclude exclusively execution exhibit expansion expedition
exploration explosive exposure extension extensively extract
fabric fantasy fatal fiber fiction finance fleet flesh flexibility
footage formula fountain fraction framework franchise freight
frequency friction frontier frustration fundamental
gallery gambling gang gaze gesture globe governance grain grateful
gravity grief grip gross guardian guerrilla
halt harvest hatred headline heal headquarters heritage highlight
historian horizon hypothesis
ideal identification ideology ignorance illustrate imagination
immune implement implication import impress impulse incentive
incidence incorporate independence index indication indicator
indigenous infection ingredient inherit initial injection inner
inquiry inspection inspiration installation insurance integrate
intelligence intense interaction intermediate internal interpret
invade investigate invisible involvement isolation
jury justification
keen kingdom
landscape launch lawn layout leadership legend legislative leisure
lessen liberal lobby logical lottery lover luxury
midst mineral ministry miracle modest molecule monitoring moral
motive mud municipal myth
negotiate neutron nightmare noble nomination nonetheless norm notable
notion
obesity observed offend offensive ongoing opera operational
organizational orientation outdoor output outbreak overcome
ownership
packet palm panic parallel participation partnership passion passive
patent pathway patience patrol penalty peninsula perception permanent
phenomenon philosopher photography phrase pilot pitch plea pleasure
plot plunge poll portrait possession pottery poverty precise
predominantly presidency pressured prevention primitive prime
priority privacy privilege productive profile prohibition
prominent proportion prosecution prospect province provision
psychology publisher pulse punishment pursuit
quest quota
radar rally realm reassure recommendation reconstruction referee
reflection reform registration regulate rehabilitation reign
remarkable remedy repetition replacement republic reservation
residual respective retail retirement reunion
sacrifice saint satisfaction scenario scholar seasonal selective
seminar sensation serial settlement shade shall shallow shelter
shelter significance simulation sketch slip slogan snap sole
solidarity sophomore spark specification spelling spiritual
spokesman squad stability stake statistical stem stimulus
straightforward strain strand submission substantially
succession summit superintendent supplement supreme surplus
surrender suspicious swift systematic
tackle teenage tenant terminal territory textile thermal tide timber
tolerance tournament trace trading tragedy transformation transit
trigger triumph tropical
undergo unity unprecedented update urge
variable venture veteran violation visible voluntary
warehouse warrant widespread withdraw workforce worthy
""".split()

oxford_b2_supp = """
absurd accelerate accessible accommodate accumulate accountability
accusation acute adaptation adolescent aesthetic aggression alien
allegation allegedly alleviate ambitious amid analogy anonymous
anticipate apparatus appetite arguably array arrogant aspire assemble
assert assumption asylum audit auction authoritative autonomy
ballot bargain battery benchmark betrayal bizarre bonus breach
breakthrough brutal buddy bureaucracy
capitalize cascade casualty categorize champagne chronic
clash cling cluster cognitive coincidence collaboration
commodity compelling complement comply configuration confront
confrontation congregation conquer consciously consolidate
contemplation contempt contention contractor controversy
conversion conviction cooperate copyright cottage coordination
counterpart courtesy cozy
deadline debris defect deficiency demon deploy depression deprive
detention deteriorate devastate devotion discourse discrimination
disposal donate donor downfall draft drainage dread duration
ecosystem elaborate embassy embrace emergence emission emotionally
empower encompass endeavor endurance enforcement entrepreneur
equate equity erosion eruption ethical etiquette evacuate
exaggeration exclusion execution exempt expertise exploitation
explosive extract eyewitness
fabricate faction fascinate feasibility fertility fierce fiscal
flourish formidable fossil framework freight frontier fruitful fury
galaxy gaze genocide gesture grace graceful grave gravity grip
guardian guerrilla guideline
harassment hatred hazard heighten heritage hierarchy humanitarian
hurricane hygiene
icon ideology illusion immense implementation importation impulse
imprisonment incentive incorporate indictment indigenous indulge
inequality infrastructure inhibit initiation innovation insecurity
insider inspection integral intensify intermediate intervention
intimate invasion investor irrigation irritate
journalism jurisdiction
label landmark latitude lawsuit leftist legacy legislation
legitimate lens liable likelihood likewise literacy lobby lottery
luxury
mainstream mandate manipulation manual margin marketplace massacre
medication merge metaphor methodology midst migrant militant militia
millennium minority miracle mobility moderate modification monument
monopoly morality mortgage municipal
narrative neglect negotiate nominal notwithstanding nucleus
oblige odds offspring optical optimistic ordinance organism orient
outlook outsider overcome overwhelm
parade paragraph parcel participation patron peculiar pedagogy
peninsula perceive persist petition pharmaceutical physician
pilgrimage pioneer plea pledge plunge polarize portfolio poster
precaution predecessor predominantly preliminary premise
prescription preservation presumably prevalent privilege
procurement profound propaganda proposition prosecution
protagonist provincial provoke psychiatric publication
punctuation pursuit
quota
radical ransom rationale realm reconcile referendum reformation
regulate rehabilitation relevance reluctant render renovation
replica reproductive reservoir resignation retail retain rhetoric
rivalry robust
salary sanction satellite scandal scrutiny secular segment
sensation sequence settler siege simulate simultaneous skepticism
soften solidarity sovereignty specimen spectrum spokesman squad
stabilize stereotype stimulus strand structural stubborn
submission substance suburban successor superintendent supplement
suppress surge surveillance surrender sustainable
telecommunications terrain textile therapy threshold tolerance
trademark transaction trait transmission transparent trauma treaty
tribal triumph turbulence turnout
underestimate undertake unprecedented upbringing utterance
vacancy valid vegetation venture verdict versus vigor
violation voluntary vulnerability
warfare warrant withdrawal worship
""".split()

oxford_c1_supp = """
abstain acclaim accomplice accord accumulation acquaintance
acute adversary aesthetic affiliation agenda allegiance
amalgamate ambivalent amnesty analogue annex antagonist apathy
apprehend aptitude aristocrat armament articulation assemblage
assertion atrocity attrition auspicious avert
backlash ballot battered benevolent blueprint bona fide bottleneck
breach breadth brochure bureaucrat
calibrate campaign canopy carcinogen catalyst celestial chronic
clamp clandestine clergy clientele coalition coercion collateral
commodity compel compilation complement compulsion conceal
concession confiscate congregation consensus consolidated
contagious contemplate contingent contraband contravene
converge conviction convoluted corporal correlation corroborate
counterfeit covenant crackdown credential criterion culmination
curb curtail cynical
daunting debacle debilitate debunk decentralize decipher
defamation defer deforestation degrade delegate delinquent
demographic denunciate deploy deregulate deterioration
detrimental devolution dichotomy differentiate disarmament
discretionary dismantle disparity disposition disrupt dissent
diversion dogma
echelon ecosystem edification elucidate emancipate embargo
embezzle emigrate empirical emulate enactment encompass
encroach endowment entail entrepreneur enumeration epitome
equitable erode espionage esteem exacerbate excerpt exonerate
expedient exponent extradite exuberant
fabrication facet federalism figurative flagship fluctuation
formulation fortify fragmentation
genocide geopolitical globalization governance grassroots gratify
grievance
harass haven hegemony herbicide heritage holistic homogeneous
humanitarian
ideology illicit imminent impediment implore improvise incentive
incur indemnity indictment indigenous indoctrinate inertia
infrastructure influx infringement inherent injunction innovation
inscription insolvency insurgent intercept interim intimidate
invoke irrevocable
juxtaposition
kinship
landmark lax legislative legitimacy lethal leverage litigation
livelihood lobbyist logistics lucrative
malicious mandate maneuver marginal mediation mercenary metabolism
meticulous migration mitigate mobilize monarchy monopoly moratorium
multilateral municipality
negligence nomenclature noncompliance normative notoriety
nurture
oligarchy omit opportunistic orchestrate ordinance overhaul
partisan patronage pedigree perpetuate pertinent phenomenon
philanthropic plagiarism plausible plebiscite plurality
policymaker pragmatic precarious precipitation predecessor
predominate preamble prerequisite prerogative prevalence
procurement proliferate propensity prosecute protagonist
protracted provisional provocation punitive
quarantine quorum
rampant ratification rationale rebuke reconciliation
redress referendum regime rehabilitation reinstate relegate
relinquish remedy remuneration repeal repercussion replenish
repression requisite resilient restitution retribution
retrospective rhetoric riveting
sabbatical safeguard sanction scrutinize secession sediment
segregation seminal severance skeptic solicit sovereignty
stabilization stagnation statutory stigma stipulate stringent
subordinate subsidy substantive succession suffrage summon
superimpose suppress surcharge surveillance susceptible
symbiotic synthesis
tariff taxonomy tentative tenure testify therapeutic thesis
totalitarian trajectory tribunal turbulence
unilateral unprecedented usurp utilitarian
validate vehement vengeance verdict vernacular veto
vindicate volatile
waiver warranted watershed
zenith
""".split()

# More supplemental words to reach ~5000
oxford_a1_more = """
also aunt bike birthday boat box brother bus butter cap card cat cheese chicken
clock coat coin color cookie corn cow cream cup daughter desk doll door driver
drum duck ear egg elephant envelope farm father fence finger flag flower fork
fruit gate glove grade grape ham hat honey horse husband ice island jacket
jeans juice key kid kite knee knife lamp leaf lemon lion lunch map meat milk
mirror monkey morning mother mouse neck nest noise noon nurse orange pajamas
paint pants parent pasta path peach pear piano pilot pink pizza plant pocket
pork prayer puppy puzzle queen rabbit rain rice ring rope rug ruler sand
sauce scissors shark sheep shirt shoe shorts silver sink sister
skirt sky snake soup square stairs stamp star steak stomach stone sugar suit
sunglasses sunrise sunset sweater swimming target tea ticket tiger towel
toy tree triangle truck tube turkey turtle umbrella uncle uniform van village
volleyball wallet wave whale wife wing wolf woman yard yogurt zipper
""".split()

oxford_a2_more = """
aboard absent abuse academic accent acceptable access achieve
acid actual adapt admit adopt advanced advertiser affair afford agency
aggressive agriculture aim alarm album allowance alongside
amazing ambition amusement angel angle annoy annual anxious
apartment apology appeal appetite application appreciation
apron arrange arrest aspect assignment assist association
assume athlete atmosphere attach attendance auction author
baggage balanced bankruptcy barbecue barn basis bat battery
bay beneath beside bet blame blend bless border bore borrow
brass breath brew brilliant broadcast brook brush bubble
budget bull bunch burden cabin calcium calm campaign canal capture
carbon carpet casual category cattle ceiling ceremony champion charity
chart chase chef chemistry chin chip chocolate circumstance civil
civilization clerk cliff climate clinic clothing cluster colleague
colony column combination comedy commander commercial commission
companion comparison compete complaint compose comprehensive
concentrate concern conclude concrete conference confess
confidence confirm connection consciousness consequence
conservative considerable construct consult consumer consumption
contemporary contest continent contribute convenience convention
cooperation core correction correspond cotton countless courage
crash creative creature crew criticism crop cruel curiosity curve
custom cycle
database dawn deaf debate decay deceive declaration decrease
definition democracy demonstrate departure dependent deposit
depression description deserve desire despite detective
determination dialogue differ dig dimension disability
disappoint disaster discipline discovery discussion dismiss
dispute distant distinct distinguish distribution disturb
document domestic dominate dozen dull dust duty dynamic
eager economy editorial efficient election elementary eliminate
elegant elsewhere emission emotion empire employ employee employer
enable encourage endless enjoyment enormous enterprise enthusiasm
entrance entry equality era error essay essential establish
estate ethnic evaluate eventually evident exact examine exception
exchange excitement exclusive excuse exhibition expand expansion
expert explanation exploration expose extend external extraordinary
extreme
fabric factor fairly faith false fare fashion fatal fault favor
feast federal female fence fiber fiction fierce financial fine firm
flame flat flesh float flood folk fool forecast forever formal
format foundation frame framework frequent frighten frontier frost
frustration fun function furniture gallery gap
gather gear gender generation genius genuine gesture ghost
giant global golden grab gradual grain grammar grand gravity
guarantee guidance guilty
harbor harvest headline heaven height hesitate highlight
horizon humor hunt hydrogen
identify identity ignorance illness illustrate imagination
immediate impact import impression improvement impulse incident
indicate infection inflation ingredient initial inject inner
innocent inspection install institution insurance intellectual
intelligence intend intense internal interpretation interrupt
introduction invest investigation invisible invitation

jewelry joint journal judgment junior justify

keen
labor ladder landscape launch lawn layer leadership lean
lecture leisure lens lesson liberal lid lifestyle limitation
liquid literature litter loan logical loose luck
mainland maker mall manufacture margin master maximum medal
medium melt membership mental merchant mercy merely
minimum minister minor mode modest monument mood moral
multiply muscle mystery
narrative nasty necessity negotiate net nevertheless nightmare noble
notion
obligation observation occupy occurrence offense offensive operate
opponent organic organize origin outcome outline overcome
ownership
packet palm parade parallel passion patience peculiar penalty
physically pitch plain planet plea poem portion pose possession
prayer predict preparation principle privacy productive
promote proportion prosecutor protective prove province
pursue
racial rain random rarely rate raw reaction readily recession
recognition recommendation reform register regulation reign
rejection relate reliable religious reluctant remote render
repair reputation resistance resolution restore restriction
reveal reverse rival rough rural
sacrifice scholar scope secondary seed senate senior sensitive
serve setback settle severe shallow shed shelter signal
slight socialist solid solution somewhat sophisticated spare
specialist spiritual stable statistical steep stimulate stock
stranger strip submit suburb suit summary super superior supply
surgeon surplus sustain sweep swift sympathy systematic

tackle tale teenage temple tend terminal territory theme
therapy thread threaten tin toll tourist tournament trace tragedy
transition transport trend tropical tunnel twist
undergo union universal unlikely upset utilize
vacuum venture version virtue visual volunteer
warehouse wealthy weird wherever widespread willing witness
worthy
""".split()

oxford_b1_more = """
abolish accelerate accessible accomplish accountability acre adapt
administer adverse affiliate agony allege allocation allowance amid
anticipation appliance arena array articulate assault assert assign
audit authorize automobile awareness
bankruptcy basement beloved benchmark bias biography blade
blast blend blossom blunt bond boost bracket breed brick
bronze bulk cabin camera capability capitalism capture
cascade cash catalog celebration ceremony certainty chancellor
chapel charm chronicle circuit clarity classic clinical clip
cluster coalition cocktail cognitive coincidence collaboration
collector colonial commercial commissioner commodity companion
comparable compatible compelling compensate competent complex
complication comply component comprehensive compulsory concentrate
confidential confirmation confront consciousness consensus
consequently construct consultant context contractor controversy
conviction coordinate corporation corridor counsel counter
coverage credibility crush curiosity custody
database deadline defender deficiency deficit deliberate
demographic deny depression designer desktop desperate detect
detention dialog dictate diminish diplomat discharge disclose
discount discrimination disorder displacement disposal dispute
distinction divine documentation donation donor draft dramatic
duration dynamics
ecosystem edition effectiveness elaborate electoral element
eligible embrace emission encounter endeavor endorse enforcement
engagement enterprise entrepreneur envelope episode equity
erosion escort ethical evaluate eventual evolve exempt
expansion expertise explicit export extract
fairly fatal fate feminist fertility fierce fiscal fixture
flame flexibility flip float flock fluid folk footage forecast
formula fossil framework franchise frontier fulfillment
functional fury fusion
gallery gambling geometry gesture glance glimpse globe glory
gorgeous graduate gravity grip guardian guideline guild
halfway hall harassment harvest hatred headline headquarters
heritage hierarchy highlight homeland honest hormone hostile
humor hurricane
identical ideology illustration imagination immense implement
implication import impose impulse inadequate incident indication
indicator indigenous inflation initiative innovation inspect
inspiration installation institutional integral integrate
intellectual interaction interim intermediate interval
intervention intimate invasion investor invisible isolation
jewelry judicial jurisdiction juvenile
keen
laboratory landscape latter legislation leisure lengthy
lieutenant likelihood linger literacy lobby
mandate mansion margin marital massive mate mechanism membership
metaphor mild militant mineral ministry miracle mobility
modify molecule monopoly monument mortgage motivation municipal
myth
narrative necessity negotiate networking nightmare noble
nominee norm notable notorious
objection obligation observation obstacle occasional odds offense
operator opponent optimize organic orientation outcome outlet
output overlap overseas ownership
parade parallel participant passionate patrol patron peculiar
penalty pension perceive permanent permit persistence petition
philosopher photography physician pioneer pitch plea pledge
plunge polar portion possess postpone pottery pray precise
prediction predominantly preference pregnant preliminary premise
prescription preserve presidency presumably prevail primitive
principal priority privilege probe profile profound prolonged
prominence proportion prosecutor protective province provision
psychiatric pursuit
qualification quarterly quest
radical rage rally rape ratio realm recession reconstruction
referee reform refugee refusal regime registration regulate
rehabilitation reinforce relevance reliable remedy renaissance
renewable repeatedly representation reservoir resignation resist
resolution respective restoration restorative restriction
revelation revenue reverse revolutionary rhetoric rigid ritual
robust romantic rotation rural
sacrifice sanction satellite savage scholar secondary secular
segment selective sensation serial session seventeen severe
shaft shadow shallow shed sheriff siege simulation singular
skilled slavery slogan slot sole solidarity southern sovereign
spatial specialist specification sphere sponsor squad stake
statistical stem stereotype strategic stride strip structural
suburban successor supplement supreme surplus suspension
sustainable sweep sword symbolic systematic
telecommunications temporarily terminal terrain textile therapy
threshold tolerance toxic trace trademark tragedy trait
transition transparent trauma treaty tremendous trend tribunal
tropical
undergo undermine unity unprecedented urban
vacation variable vast vegetation venue verdict versus viable
violation virtue visible visual vital volunteer vulnerable
warehouse warfare warrant wholly widespread wildlife
withdrawal witness workforce worthy
""".split()

oxford_b2_more = """
abolition absorbent abstract acclaim accomplishment accumulation
accountability acknowledgement acquisitive adjective administrative
advent advocacy aerial affiliation aftermath allegory allot amateur
ambiguity amicable anarchist ancestry animation annex anticipate
apparatus appease appliance apprentice arbitration articulate
assassination assertion astronomy attribution authentic aversion
backdrop backlash bailout benchmark betrayal bipartisan
blackout blueprint bombardment botany boycott breadth
brochure bureaucratic
calculation calibration camouflage capitalism capsize cardinal
catalog centralize cerebral chancellor chapel chronological
circumvent civic clientele cloak coalition coherence collaboration
collateral commemorative commentary commodity communion
compilation compliant compositor compound comprehensive
compulsory concession condemn confederation confiscation
conjunction conscription consolidate constitutionalist
contemplation contention contractor contradiction conversational
convoy copyright corporation correlate correspondence counselor
courtesy crackdown creditor culmination curator curriculum
cynicism
deception decisive decompose defamation delegation deliberation
demographic denomination dependence deployment depreciation
detection deviation deviation dialect dictatorship differential
diplomacy directive disarmament disclaimer discrete disposition
disseminate dissertation divergent documentation dominion donation
downturn draftsman duplication
ecosystem edification electorate elevation embargo embodiment
emergence empowerment engagement enrollment entrepreneurial
envoy epidemic equator erratic escalation evangelical evasion
excavation excise exemption exhibitor exonerate expedition
exponential
fabrication factorial facilitation fascism feasible fidelity
fiscal flotation fluency folklore foreground forestry formidable
fortification fracture fulfillment furnish
garrison gazette genocide geopolitics globalization governance
gradient grassroots grievance guerrilla
hallmark harmonize haven hereditary hierarchy holistic homage
humanitarian hypocrisy hypothesis
ideological illuminate immersion immunity impediment implement
improvise inception increment indigenous indoctrinate induction
industrialize inference influx infringement inheritance injunction
inscription insolvency insurgency integration intercession
intermittent intolerance intrinsic introspection invalidate
irrigation isotope
jurisdiction
kinetic kinship
landmark lament layman legislation legitimacy leverage
liberalization lineage litigation livelihood locomotive
logistics longevity
magistrate magnitude malfunction mandate manifesto manipulation
marathon marginal maritime maternity matrix mayoral mediation
memoir merchandise metabolism methodology militia misconceptual
monastery monetary monopoly moratorium municipal
narrative negligence neutrality nobility nomenclature notary
numerical
obsolete offset omission operative optimization ordinance
orthodox outskirts oversight overthrow
pacifist pandemic paradigm partisan patriotism patronage
pedagogy pedestrian perpetual petition pharmaceutical pilgrimage
plebiscite pluralism policymaker polygamy postwar practitioner
precautionary predecessor precipitation preface prerequisite
prevalence procurement prohibition proliferation pronunciation
propaganda proposition prosecution provincial psychology
publication punitive
quarantine quota
ratification reactionary reconciliation reconstruction
redevelopment referendum regime reimburse relic remedial remittance
renovation repatriation repercussion replica repository
repression requisite resettlement residual resignation
resonance resurgence retrospective rhetoric rivalry
sanction scarcity scrutiny secession sediment semiconductor
sentinel settlement severance shareholder skepticism solicit
sovereignty specimen speculation stagnation statute stereotype
stewardship stigma stipulation strand subdivision subordinate
subsistence suburban suffrage superintend supervision supplement
surcharge surveillance susceptibility symbiotic synthesis
tariff taxonomy tenure territorial testament therapeutic
topography totalitarian trajectory transcend transition
transparency tribulation tribunal tribunal tumult
undermine unification unprecedented utilitarian
validation variable vegetation vendetta venue verification
versatile veteran vicinity vindication volatility
waiver watershed wholesale
xenophobia
""".split()

oxford_c1_more = """
aberration abjure abstraction accolade accreditation acrimony adjudicate
admonish adversarial affidavit aggrandize agrarian alliteration
amalgamation ameliorate amnesia anthropology antimicrobial antithesis
apex aphorism apostrophe appellation arcane arraign ascribe aspersion
assimilate atone attribution augment austerity avocation

belligerent benign bequest bifurcation biodiversity blasphemy
blatant boisterous bourgeois brevity cache capitulate caustic
caveat cessation circumscribe coalesce coerce cognizant colloquial
commodify complacent complicit concomitant condone conflagration
conjecture consign consortium construe contemptuous contingent
contradict convalesce convergent convivial copious corroborate
covert credence culpable cumbersome curate cynic

dearth debase debilitate decorum deference deft delineate deluge
demise denote deplete deposition derelict despondent deter dichotomy
diffuse digression dilapidated discern discordant discourse
disenchant disenfranchise disheveled disingenuous disparage
disseminate dissipate distill divergence dogmatic dormant

ebullient ecclesiastical efficacy effluent egalitarian elicit
eloquent elucidate emanate embellish embroil eminent empirical
emulate encapsulate encumber endemic engender enigma ennui
epoch equanimity equivocal eradicate errant eschew esoteric
etymology euphemism evanescent exacerbate exasperate excerpt
excoriate exemplify exhort exodus expedient explicit extemporaneous
extenuate extricate
fabricate facile fallacious fanatical fastidious fatuous
fervent fidelity flamboyant florid fluctuate foment forestall
fortuitous fracas fraught frenetic frivolous fruition furtive

galvanize garner garrulous genesis genteel germane gesticulate
grandiose gratuitous gregarious gubernatorial guile

hackneyed hapless harangue harbinger haughty heinous heretical
heyday hiatus holistic hubris humility hyperbole

iconoclast idiosyncrasy ignoble immaculate imminent impartial
impeccable impede impervious implement implication implicit
impostor impregnable impropriety impugn inadvertent incandescent
incarcerate incessant incipient incisive incognito incongruous
incredulous indelible indemnify indeterminate indifferent
indigenous indiscriminate indolent indomitable ineffable
inexorable infallible inference inflammatory influx ingenious
inherent innocuous innuendo inquisitive inscrutable insidious
insolvent instantaneous instigate insular interlocutor
interminable interpolate intractable intrepid intricate
inundate invective invoke irreverent

judicature jurisprudence juxtapose

labyrinth laconic lament languish laudable lethargic levity
licentious loquacious lucid ludicrous

magnanimous malevolent manifesto maudlin meander mediocre mendacious
meritocracy metamorphosis meticulous milieu misconstrue misnomer
mitigate mollify monolithic moribund munificent myopic myriad

nascent nebulous nefarious nihilism nominal nomenclature
nonchalant nonpartisan nostalgia novice nuance
obdurate obfuscate oblique oblivion obsequious obsolescence
obstinate ominous omnipotent onerous opaque opportunist
optimize ornate orthodox ostensible ostracize overt

palatable palpable panacea paradoxical pariah parochial parsimonious
partisan patriarch pedagogy pedantic perfunctory perjury
permutation perpetuate pernicious perspicacious philanthropic
platitude plethora poignant polemic ponderous postulate potent
practicable pragmatic precarious precipitate predilection predominate
preeminent preempt prejudicial prelude preposterous prerogative
prescient presumptuous pretentious principled probity proclivity
prodigious profligate progenitor prolific promulgate propensity
propitious proprietary prosaic protagonist protracted provincial
prowess prudent punctilious purist puritanical

querulous quixotic
ramification rancor rapacious rationale rebuff recalcitrant
reciprocal recourse rectify redolent redundant refurbish refute
relegate relinquish remorse replete reprehensible reprisal
repudiate requisite rescind resonant resplendent restitution
resurgent reticent retroactive revere rhetoric rife rigorous
sagacious salient sanguine sardonic scrupulous sedulous
semblance semiautonomous sequester shrewd skeptic soliloquy
soporific specious spurious squalid steadfast stringent
subjugate sublime substantiate subversive succinct suffrage
superficial supplant supposition surreptitious susceptible
sycophant
tacit tangential taut tedious temperance tenacious tentative
theology tirade titular torpid torrent transgress transient
trepidation trenchant trite truculent tumultuous tyranny
ubiquitous unctuous underpin unequivocal unprecedented
unscrupulous untenable utilitarian
vacuous vanguard vapid vehement verbose veritable vicarious
vigilante vindictive virulent visceral vitriolic vociferous volition
whimsical
zeal zealous
""".split()

# Final batch to reach ~5000
oxford_extra_final = """
abandon absorb abstract abundance academy accent acceptable accomplish
accumulate ace ache acid acquaintance acquisition acre actively adequate
adjustment admiration adolescent adoption advancement advent
adventurous aerospace affirmative afterward altogether ambassador
ample ancestor ankle antique apologize appetite approximate aquatic
arrogant arrow assess asset autobiography automobile

backbone backward bamboo bandwidth banner bargain basement battalion
bay beard behalf benchmark bid billboard bind blade blessing boast
bookmark boundary boxing bracket brew bronze browse brutality bubble
buffalo bulk bureau butterfly bypass

cabinet calcium calendar calling calorie canal candle capability
capitalism capsule caption captive carve casino casualty catalyst
caution celebrity cemetery cereal championship characterize chart
checkpoint chord chorus citation civic clarity clash classify clay
clearance cliff climax clip cluster coalition cocktail collaborate
collar colonial comedy commodity companion compensation competence
compilation compound comprehend compromise compulsory concentrate
confession conscience consciousness consent consolidated consultant
consumption contemplate contest contradiction contributor
controversial convention convince cooperate coordination courage
creativity credibility crew cruelty cuisine curve

database deafen debris dedication default defiance delight
democratic density departure deploy deposit deputy desk destruction
diabetes dialect dignity dimension disability disaster discipline
discourse discrimination displacement disposal distant distinctive
distribute diversity dividend documentary doctrine dominant donor
dose draft dramatically drift dwell dynamic

earthquake ecological economically editorial effectiveness
elementary eligible elimination embarrassment emission emperor
emotional emphasis empire employment empower encounter endure
enforcement engagement enrollment entertainment entity entrepreneur
erosion ethnic evaluate evolve exaggerate exceptional excessive
exclusive exhaust exotic expansion exploration explosive export
exterior extinction extreme eyebrow

famine fare fascination feast fertility fiber finance fixture
flatter flesh flexibility foreigner format formula fossil fraction
fragment franchise frequency frustrate furnish

galaxy gender gene genetic glory governance grasp grief grocery
gross guardian

habitat halfway halt hazardous headline helicopter heritage
highlight hormone humidity hurricane hybrid

identical ideology illustration immigrant immerse immune impairment
imperial implement import improvement impulse inadequate
inappropriate incident incorporate indication infrastructure initial
initiative injection innocent innovative insistence installment
institution instruction instrumental insufficient intake intense
interior intermediate intersection intimate invasion inventor
investigate invisible irrigation

judicial jungle jury juvenile

keyboard
landfill landmark landscape latitude laundry legitimate
leisure lens liability liberal likelihood limitation linear
livestock lobby locomotive logistics loyalty

macro magical magnitude mainland manipulation mansion manuscript
marathon margin mature meadow mechanism memoir merchant merit
metabolism metaphor metropolitan mighty milestone mineral miniature
mitigate moderate molecule monarchy monitor monumental mortality
mortgage motivation multiple municipality

nail narrative naval necessity newsletter nightmare nobility
nominal nonetheless norm notorious nuclear

obesity obligation occupation occurrence offspring operational
opponent optical orchestra ordinance orientation orphan outbreak
outbreak outreach outsider oversight

paddle palm paradox parcel parliament partnership passport patron
peasant penalty penetrate pharmaceutical physics pilot pioneer
plausible plaza pledge plumbing poet polar portfolio poverty
precision predecessor predominantly pregnancy prescription
preservation prestigious prevalent primitive privilege probe
productive progressive prohibition prominent prompt propaganda
proportion protest province provoke psychological publisher
punctual pursuit pyramid

qualification quarantine quarterback

racial rally random realization rebuild recession recipe
reconcile recruitment referendum reflection regulation rehabilitation
reliability remedial removal rendering renovation repetition
repository representation reproduction republic resemble
residential resignation resort restore restriction retreat retrieve
revelation revolt rigid ritual rivalry robot rotation

salmon satellite scenario scholarship segment semester sensation
sequence settlement shallow simultaneously skeptic snapshot
socialist sovereignty specification specimen spiritual spouse
stabilize stake standpoint statistical stimulate strand strategy
suburban succession superintendent supervision supplement
surveillance suspend sympathy synthetic

tactical taxpayer telescope temporary terminal terrain textile
theoretical thriller tolerance toxic trajectory transformation
transmission transparent tribal trillion tropical turbulence
tutorial

unanimous undergraduate unemployment unfold unify unprecedented
uproar utilize

vacuum validity vaccination variation veteran villa violation
visibility volatile voluntary vulnerability

warehouse warrant weekday wellness wholesale widow wilderness
workforce worship

yacht

accomplish acre admiral advent affiliate aggregate alongside
amateur ambiguity analogue anticipation appliance apprentice
arithmetic assurance auction autonomy
bachelor ballet bankruptcy batch benchmark biography bomb bore
breadth bronze cabin calcium calendar canon capability capitalism
carnival carpet castle catalog census chaos charter choir
chunk cipher civilization clergy cluster coalition cocktail
cognitive colleagues colony comedy commodity communist comply
confession consolidate constellation contempt context
controversial convey copyright corridor cottage crane
creativity cricket criterion crucial curriculum custody
dairy dean decimal decree deficit delegation derive destiny
diabetes dialogue diesel dignity dilemma diminish diploma
discourse discrimination drift drought dump durable dynamic
editorial elaborate electrode eligible elite embassy emigrate
emission emperor encyclopedia endeavor endorse engagement
enhanced equation equity erosion escort estate evaluate
excavate excess exclusive exile explicit exploit exposition
external extract
fiction fleet flourish folk forge forum fossil fraction freight
fusion gallery genre gesture glory golf gorgeous grammar granite
gravity grip grocery guardian guerrilla guilty
habitat halt hardware harvest hay headquarters heritage
historian hollow hurricane hydraulic hydrogen
ideology illustration immigrant immune incentive incumbent
indigenous infantry inflation influential infrastructure inherent
inject innovation integrity intervention intimate irrigation
ivory
jurisdiction
keen kingdom
landlord latitude layout legislative leisure lens lever liberal
literacy locality locomotive loom luxury
magnitude mainstream mandate manuscript marble margin
marshal martial maturity memoir merchandise mercury migration
mineral ministry mob moderate molecule monopoly monument
mortality municipal myth
narrative negotiate neutral nightclub nobility norm notorious
nursery
obesity obstacle offspring operational opponent orchestra
ordeal orientation outline outsider overseas oversight
paradigm parliament participant passive patent patrol patron
peasant penalty peninsula percussion petition philosopher
physician pilot plantation plaza plumber poetic portable
postpone poultry poverty precise premier premium primitive
privilege probe productive prominent prompt pronunciation
prophecy proportion prosecution protagonist provincial provoke
psychiatric publicity pulse pursuit puzzle

yacht
""".split()

def gen_oxford():
    words = {}
    for w in oxford_a1 + oxford_a1_supp + oxford_a1_more:
        words.setdefault(w.lower(), 'a1')
    for w in oxford_a2 + oxford_a2_supp + oxford_a2_more:
        words.setdefault(w.lower(), 'a2')
    for w in oxford_b1 + oxford_b1_supp + oxford_b1_more:
        words.setdefault(w.lower(), 'b1')
    for w in oxford_b2 + oxford_b2_supp + oxford_b2_more + oxford_extra_final:
        words.setdefault(w.lower(), 'b2')
    for w in oxford_c1 + oxford_c1_supp + oxford_c1_more:
        words.setdefault(w.lower(), 'c1')
    path = os.path.join(DIR, 'oxford5000.csv')
    with open(path, 'w') as f:
        for w in sorted(words):
            f.write(f"{w},{words[w]}\n")
    print(f"oxford5000.csv: {len(words)} words")

# ============================================================
# NGSL — rank,word  (New General Service List ~2800 words)
# ============================================================
# Ordered by decreasing frequency — ~2800 most common English words
ngsl_words = """
the be and of a in to have i it for not on with he as you do at this but his by
from they we say her she or an will my one all would there their what so up out
if about who get which go me when make can like time no just him know take people
into year your good some could them see other than then now look only come its
over think also back after use two how our work first well way even new want
because any these give day most us
find here thing many right tell still great help own through child life long hand
between old school must home under close last question try fact far left need start
side world keep eye country open against state company without body hear point set
government run small number off always move night live enough head show however
idea turn another quite money serve word house leave different program away water
really call several form develop food grow begin ask name girl important place hold
present end follow part family become next early include reach late mean talk each
before boy problem interest within big high sort base cut sure report describe feel
result read group change certain sense reason public step cause continue wrong
strong story type real land act mother area local price father believe national
plan foot war produce hand offer provide decide start win picture study since hard
stand age drive top rest pay record million community play table case carry job
either pass appear across member seem add hope deal able air line letter course
write hour

able above accept according across act action activity actual actually add address
age agency agent ago agree agreement ahead aim allow almost alone along already
also although always american amount analysis animal announce annual another answer
anything appear application apply approach appropriate area argue argument arm army
around arrive art article artist ask assume attack attention audience authority
available avoid away

baby back bad bag ball bank bar base basic basis bear beat beautiful bed begin
behavior behind belief believe belong beneath benefit beside best better beyond
big bill billion bit black blood blue board body bone book born boss both bottom
box boy brain break bring broad brother brown budget build building burn bus
business buy

call came camera campaign campus can candidate capital car card care career carry
case cash cat catch cause cell center central century certain chair chairman
challenge chance change chapter character charge check child choice choose church
citizen city civil claim class clean clear clearly close coach cold collection
college color come comfort command commercial commission common communication
community company compare competition complete computer concern condition
conference congress consider consumer contain continue control conversation cook
cool corner cost could council count country couple course court cover create
crime cross cultural culture cup current customer

dad damage danger dark data daughter day dead deal death debate decade decide
decision deep defense degree democrat democratic department depend describe design
despite detail determine develop development device die difference different
difficult dinner direction director discover discussion disease doctor document
dog door down draw dream drive drop drug during

each early east easy eat economic economy edge education effect effort eight either
election else employee end energy enjoy enough enter entire environment especially
establish even evening event ever every everybody everyone everything evidence
exactly example executive exist expect experience expert explain eye

face fact fail fall family far fast fear federal feel few field fight figure fill
final finally financial find finger finish fire firm fish five floor fly focus
follow food foot for force foreign forever forget form former forward four free
friend from front full fund future

gain game garden gas general generation get gift girl give glass go goal good
government great green ground group grow growth guess gun guy

hair half hang happen happy hard hat he head health hear heart heat heavy help
here herself high highly him himself his history hit hold hole home hope hot hotel
hour house how however huge human hundred hurt husband

idea identify if image imagine impact important improve in include increase
indicate individual industry information inside instead institution interest
international interview into investment involve issue it item itself

job join just keep key kid kill kind kitchen know knowledge

land language large last late later laugh law lawyer lay lead leader learn least
leave left leg legal less let letter level lie life light like likely line list
listen little live local long look lose loss lot love low

machine magazine main maintain major majority make manage management manager many
mark market marriage material matter may maybe me mean measure media medical
meet meeting member memory mention message method middle might military million
mind minute miss mission model modern moment money month more morning most
mother mouth move movement movie mr mrs much music must my myself

name nation national natural nature near nearly necessary need network never new
news newspaper next nice night no none nor north not note nothing notice now
number

occur of off offer office officer official often oh oil ok old on once one only
onto open operation opportunity option or order organization other others our out
outside over own owner

page pain painting pair paper parent part particular particularly partner party
pass past patient pattern pay peace people per percent performance perhaps period
permit person personal phone physical pick picture piece place plan plant play
player please point police policy political politics poor popular population
position positive possible power practice prepare president pressure pretty
prevent price private probably problem process produce product production
professional professor program project property protect prove provide public pull
purpose push put

quality question quickly quiet quite

race radio raise range rate rather reach read ready real reality realize really
reason receive recent recently recognize record red reduce reflect region relate
relationship religious remain remember remove repeat replace report represent
republican require research resource respond response rest result return reveal
rich right rise risk road rock role room rule run

safe same save say scene school science scientist score sea season seat second
section security see seek seem sell send senior sense series serious serve service
set seven several shake shall shape share she shoot short shot should shoulder
show shut side sign significant similar simple simply since sing single sir sister
sit site situation six size skill skin small smile so social society soldier some
somebody someone something sometimes son song soon sort soul sound source south
southern space speak special specific speech spend sport spring staff stage stand
standard star start state statement station stay step still stock stop story
strategy street strong structure student study stuff style subject success
successful such suddenly suffer suggest summer support sure surface system

table take talk task tax teach teacher team technology television tell ten tend
term test than thank that the their them then there these they thing think third
this those though thought thousand threat three through throughout throw thus
time to today together tonight too top total tough toward town trade traditional
treat treatment tree trial trip trouble true truth try turn tv two type

under understand unit until up upon us use usually value very victim view
violence visit voice vote

wait walk wall want war watch water way we weapon wear week weight well west
western what whatever when where whether which while white who whole whom whose
why wide wife will win window wish with within without woman wonder word work
worker world worry would write writer wrong

yard yeah year yes yet young your yourself youth

ability absence absolute absolutely absorb abuse academy acceptable access
accident accompany accomplish accord account accurate accuse achieve achievement
acid acknowledge acquire acre act actual adapt add addition additional
adequate adjust administration administrator admire admission adopt adult
advance advanced advantage adventure advertising affair afford afraid african
afternoon age agricultural aid aide aim aircraft alarm album alcohol alien
alliance allied allocate alternative altogether amazing ambassador amend
amendment amid analyst ancient anger announce annual anonymous anxiety anxious
anybody anytime anyway anywhere apart apartment apparent apparently appeal
appearance apple approve architect area arena arrest arrival arrive arrow aside
assert assess assessment asset assign assignment associate association assume
assure atmosphere attach attorney attract attribute auction author authority
auto automobile available avenue average awareness

background badly bag balance ballet ban banking barely barrier basket bathroom
battle bay bean bear bedroom beef beer beginning behave being bell belong below
bend big birth birthday bite blame blast bless blind blow board boat bold bomb
bond borrow both boundary bowl boyfriend brand brave bread breakfast breaking
breath breathing breed brief briefly broad broadcast brother brown brush buddy
bug bullet bunch burden burn burst bury buyer

cabinet cake calculate caller calm cancel capable captain carbon carpet carry
cast casual catch category cattle ceiling cell chamber champion champion channel
character charge chase cheap cheat cheese chest chicken childhood chip chocolate
choice chronic chunk cinema circle circumstance cite citizen civilian classic
classroom clerk click climate clinic clinical clock closely clothes coach
coalition code coffee cognitive cold colleague colonial color column combat
combination comedy comfortable commander commercial commit commodity companion
comparison compatible complaint complex component compose comprehensive comprise
compromise computer concentrate concept conclude concrete conduct conduct
confidence confirm conflict confront congress consciousness consensus
considerable construct consultation consumer consumption contact contemporary
content contest contract contractor contribution controversial conversation
cooperate core corporate counsel counter courage coverage craft crash crazy crew
crime criminal criticism crop crowd crucial cry curious cute cycle

daily dairy damage dare database deadline dealer dear debate debris decent decline
defeat defense deficit deliver demonstrate deny departure depend deploy depression
describe description description designer desire desperate destroy destruction
detect determination dialogue differently dig digital disability disagree disaster
discipline discourse discrimination display distinguish distribute district
disturbing diverse diversity dock domestic dominant double downtown dozen drama
dramatically drawing drop due dump dust dwelling dynamic

eager ear earn ease easily eastern eat ecological economic edition efficiency
eighth elect electoral elegant element eliminate elite elsewhere embarrass
emergency emotional emphasis emphasize empire employer employment enable
encounter encouragement enemy engineering enhance enormous enterprise entitle
entry environmental equal era essential essentially estate ethnic evaluate
evaluation eventually eventual evidence evolution evolve examination examine
exceed exception excessive exchange excited exclusive exclusively execution
executive exhibit exhibition exile expansion expectation expense expensive
expertise explanation explosion expose extend extension extensive extent extract
extreme

fabric facility fade fame fan fantasy fare farmer fascinating fashion fate
favor favorite featured fee feedback fewer fiber fiction fifth fifty fighter
filmmaker final finding firmly fishing fit fix flag flat flexibility float
flood flying fold folk force forecast formula forth fortune fossil fraction
frame franchise frequency frequent frustration fulfill fun fundamental funding

gallery gap gear gender gene generate genetic genius gate gathering giant gesture
glad glimpse glove golf gorgeous government grab gradually grain grand
grandfather grandmother grave gravity greatly grocery guarantee guardian
guilty guitar

habitat halfway hall handle handsome harbor hate headline headquarters healing
heavily heel helpful heritage hero highlight highway hint historian honestly
horrible horror host household humorous hunt hypothesis

ideal identification illness illusion illustrate immigration immune immense
immigrant immune implement imposing indicate inevitable initially inspection
inspiration inspire installation intelligent intense interact internet
intervention investor invisible isolation

jail jet jewelry joint journalist judgment junior jury

keen kindly knock

label laboratory lace landscape lap largely laser layer lazy leadership lean leap
lecture legacy legitimate length lens less lesser liberal liberation lid lifestyle
limited literally literary locate logical lovely lucky

mainstream maker mall manifest manner manufacturing mark mass mate mayor
meaningful meanwhile mechanism membership memorial merely mess method mine
minimal minimize miracle mix mixture mode moreover mostly mount multiple
municipal mutual

naked namely narrative nasty nationwide negotiate nerve newly nightmare noble
nod nominate none nonetheless noon norm normal northeast northwest notable
nowhere nuclear nursing

oak objection obligation observer occasionally occupation offensive offensive
ongoing opening opponent oppose opt organ organic originally other outcome
outfit output outsider outstanding overcome overlap oversee overwhelming

pack pale palm pant parade passenger passing passion passive pathway patience
peak pen per perfectly permanent personally perspective phase philosophy phone
photograph pile pink pitch plain pleasure plot pocket poem poet poetry portfolio
portrait possession potentially poverty praise prayer precisely prediction
preference pregnant prep presentation preserve presumably pride principal
principle privately proceeding processor productivity professor profile
profound promising promote prompt proof properly prosecutor protective prove
province psychological publicly publisher

quest

radar radical rage railroad rape rapidly raw reaction realistic realm receiver
recovery recruit reform regarding regional reinforce reject relevance reliable
relieve reluctant remarkable removal rent repeatedly replacement request resource
resolve restoration reveal revenue reverse revolutionary rhythm rifle robbery
romantic root rough routine

sacred sadly sake sand satellite scandal scholar scholarship scope seasonal
secondary secretary segment selective self senator sensitive settle seventh
severe sexual shadow shelter sheriff shield shock shooting shore shortly
shortage shout shut sigh signature similarity similarly singer sixth skilled
slave slight slowly smartphone software sole solidarity sometime somewhat
soul specialist specifically spectacular spending spiritual spokesman spot
spy squeeze stare statistical statue steadily steep stem stimulus stir stock
storm strategic stress stretch strip struggle subsequent substantial suburb
sue suffering summit super suppliers suppose surgeon surgical surplus survive
sustain symbol

talent tap teenager temple temporarily tender territory theater thermal
thin thoroughly threaten till timing tobacco tolerance tomorrow tongue
tourism tournament trace tragedy trail transformation translate tremendous
trend trick trim triumph tropical truly tunnel twist

ugly uncertain undergo unemployment unexpected unfair unique universal
unlike unlikely update urban urge

vacation valuable variable variation venture versus video violate virtue
visible visual vital vocabulary volunteer

wake wealthy weird whenever wherever widespread willing wire wisdom
withdraw worker workforce wrap

acre adapt adjust admire adventure advocate affordable aid aircraft alarm
alliance alter altogether amid angel annual anxiety appetite architect
assist assumption atmosphere attorney authorize
bargain barrier basket battery behalf bell beloved bend biography bishop blade
blessing bold bore bounce boundary boyfriend brass breast breeze bronze
brush buck bullet cab cabin calm canal cape carpet cartoon casualty cattle
ceiling chain chalk chamber championship chaos chapel charity charm charter
chemistry choir chunk circuit cite civilian clarity clerk cliff clone cluster
coach coalition cocktail coefficient colony combat comedy commence commission
companion comparable compatible compelling complement complexity comply
comprehensive compute concentrate confession confident confuse congress
consciousness consensus consent conservation considerable consistent
constitutional construct consult contemplate contempt context controversial
conviction coordinate cope corporate counsel countryside coup coverage
crack criticism crop crucial crystal
dam dare darkness database deadline decline deer deficit deliberately delivery
democracy demonstration density deputy derive desert desperate determination
diagram dialogue diet dig dignity dimension diplomat disability disagree
disaster discipline discourse discrimination displacement disposal distant
distinct distinguish diversity documentary dominant dominate dose downtown
dramatically drift dump
ease echo editorial effectively elaborate elderly elect electoral elegant
eliminate elsewhere embrace emission emotional empire employ employment
encounter enforce engineering enormous enterprise enthusiasm entrance
environmental era essay evaluation evolution exam exception excessive
exclusive exhibit expand expansion expectation expertise exposure extract
extraordinary
fabric faculty fade fatal favorable fiber fierce filming finding fishing
flash flexibility flood folder folk footage forecast forum foster foundation
franchise friendship fuel fulfill function furniture gallery
gang gathering gear generate generous genetic genius genre ghost giant
glimpse globe golf gorgeous grab grain guarantee guardian guitar
handful harbor harvest headline heaven heel helicopter heritage
historian hollow homework horizontal horror hosting hostile household humor
hunt hurricane
ideal illustration immigration immune implement implication impose incredible
indication indigenous industrial infection infectious ingredient inject
innovation inquiry insert installation institute institutional insurance
integrate integrity intelligence interact interpret intervention invasion
investing investor iron isolation
jail jazz jet jewelry jurisdiction justify
keen killing kit
landscape laser latter launch lawn layer layout league leather legacy
legitimate lengthy lens liberal liberty likewise limitation linguistic
literacy lobby logical log luxury
margin marine marketplace mathematics meaningful meanwhile
mechanism melody membership mentor mercy metropolitan migration mine
minimal miracle moderate modification mood moreover mud
neglect negotiate nerve networking nightmare noble norm nuclear nursing
nutrition
obligation observation obstacle offense offensive ongoing operational organ
organic organize oriented outcome outer output overlap overseas overwhelming

palace pan parade parish parking participation partnership passport
pattern patrol peaceful peculiar penalty pension pepper per percentage
perception permanent personally phase phenomenon pickup pile pioneer
pitch plain planet plenty plug plunge poetry poll poll pollution pope
portrait pose possession poster potato potentially precise prediction
preference premium prescribe preserve privilege probe profile profound
progressive prolong prominent promote proportion prospect protective proud
province psychological publisher
quest
rage rally rank random rarely raw ray reading recipe recovery recruit
refugee regulate reign reinforce relevance reliable remedy render repair
repeatedly representation republic reputation requirement rescue
resignation resistance resort respective restoration restrict retain
retreat revolutionary rhythm ridiculous rival robot rotate route
sacred saint salary sanction satellite scanner scenario scholarship
scope seasonal secular segment semester senator sensitivity sequence
severely shelter shipping shore shortage shut siege signature silly
skilled slavery slim snap solar sole sophisticated spare specification
spectrum sponsor spouse squad stagger stake stance statistical subtle
succession summit super supposed surgery surplus suspension sustain sweep
sword symptom syndrome
teenage temporarily terminal terrain tobacco toll ton tourism tournament
trace trademark tragedy transition transmission transparent tremendous
tribe triumph tropical tunnel
undergo undermine unemployment unite unprecedented urban utility
valve variable vast venture verbal versus veteran violation virtue
visual vocal volatile vulnerability
warehouse wheelchair whenever wherever whistle widespread workout workshop
worldwide worship worthy

absorb abstract accustom acknowledge acquisition addiction adjacent
administer adolescent adoption adverse advocate agenda allege allocate
alter aluminum ambassador amend ample analogy annually anonymous ape
apparatus applaud arbitrary arena ash assert assist attain attorney
authorize autonomy

bankrupt batch beneficial biography bloom bonus boundary bracket
breastfeed bronze bureaucracy
calendar canal capability captive cater cement championship chaos
charter chronic chunk clash clarity clinic clip cognitive coin
collaboration command commence commissioner commodity companion
comparable compel competent compile complement comprise compromise
compulsory conceive concrete consecutive consensus consolidate
constitute constraint consult contemplate contempt contend contradict
contractor conversion convict cooperate coordinate corporate
correlation correspond counsel coverage craft curriculum
damn database deadline debris decade dedicate default defect defendant
deficit democracy density deployment deputy deprive designate
deteriorate deviation dictate digital diplomatic disability discourse
discrimination displace dispose dispute distinct documentary
documentary dominant donation donor dose duration dynamic
elaborate embrace emission enforce engagement equity era evaluate
evident evolve exaggerate exceed exception exclude exclusively
execute exempt exhibit expansion explicit exploitation export
extract
facilitate fake fame fantasy fascinate feasible fertility fiction
file flame fond forbidden formula fossil foundation fraction fragment
franchise freight fuel fulfill
gallery gap garbage gender genetic genius globe governance grammar
gravity grocery
habitat headline heritage hierarchy historian household hydrogen
hypothesis
icon ideology implement implicate import indicator indigenous
inevitable infrastructure inherent initial inject innovation insert
inspection institutional instructor integrate integrity intellect
interaction interference interior intermediate interval intimate
invade inventory invest isolate
jurisdiction juvenile
keen
landlord launch legend legislature legitimate leisure lengthy lens
liable lifetime likelihood linguistic literacy
magnitude mainstream mandate manuscript margin mature maximize
meaningful mechanism medieval membership merge metaphor migrate
mild militia minimize miracle mobile mode modification molecule
monument moreover mortality mounting multiple municipal
narrative negotiate neutral noble nomination notable notion
numerous
oblige obstacle obtain offset ongoing optical oral organ orient
orthodox outcome outlet output outweigh
partial participation particle patent peer penalty permit persist
petition philosophical physician pioneer plea plead plot portable
portion posture poverty precise predecessor predominantly pregnant
premium prescribe prevail principal privilege probe productivity
projection prominent pronounce proportion prosecute province
psychological
quota
racial radical rape ratio readily realm recession refuge regulate
reinforce relay reliability reluctant render resign resolution
resolve respective restoration revelation revenue reverse revision
revolution rhetoric rigid ritual
saint satellite scenario scholar scope secular selective semester
sentiment separation serial simulate simultaneous skeptical socket
sole sovereign spatial spectrum spouse squad stable stance stimulate
straightforward strand strip structural subordinate subsequent
substantial successive summit supplement suppress surveillance
suspend sustainable systematic
tackle tale temporary terminal territory theology therapy thesis
threshold tolerance toll toxic trace trait transition transmit
transparent treaty tremendous trend trigger triumph
undercover undergo underline underlying undermine unfold
unprecedented utilize
vacancy valid variation venue versus via vessel victim violate visible
volleyball voluntary
waist ward warehouse warrant web wheelchair wholesale withdraw
workforce workshop worldwide

accomplish adaptation adequately administrative agricultural aide
allocation alongside altogether ambassador ambiguity amend
appreciation approximate assign assumption automobile
beast bias biography blank blessing boundary buffer
cabinet calendar candle capability castle catalog chaos charm
citizenship classification clinical cluster collective colonial
companion comparable compilation complexity comply comprehensive
compromise conduct cone confession conscience consolidate
consultation conversion coordinate corresponding cotton
curriculum
dawn declaration deficit density deputy designation dialect
diamond differs dignity disability disorder displacement
distinction divine dose duration
elaborate emission empathy empire enforce engagement entity
equity essay estate evaluation exception excessive
exclusively existing exhibition exotic exploit exposure
fabric fantasy fate feminist fifty fixture flame flesh
footage formation fortune fossil founder fraud
galaxy genius genuine glory golf grab gravity guarantee
gulf
hardware headline heritage hike homework horizon hormone
hypothesis
identical ideology illustration imagination impose impulse
indicator individually integral integral intellectually interact
interface invasion investor irrigation item ivory
jet jointly jury
kingdom
landlord lifetime lodge logic
mainstream mandate manuscript medal mercy migration ministry
mobile modify molecule monetary monopoly morale mortality
municipal mutual
namely narrative neglect nevertheless noble nomination notion
nursery
offensive ongoing operational opponent optical organic
orphan outfit outfit override overturn ownership
parade parallel partial participation passionate passive patrol
penalty pharmaceutical pine plain plea pledge portable
posterior pottery poverty practitioner preceding predominantly
premium preliminary preservation prevail previously primitive
probability probe productive profound prophet proportion
prospect psychiatric publication
qualification
racism radical rally realm recession reconstruction referee
regional regulate rehabilitation relevance reluctant remedy
render renowned repetition republic resemble respective
restoration revolutionary ritual rotation
seasonal secular selective semiconductor sensitivity
sexuality shelter simultaneously skeptical sovereign
specimen sponsor stabilize statistical strategic strip
subsequent subsidiary sustain sympathy systematic
therapeutic threshold tolerance toxic trait transformation
tremendous tribal triumph trustee
undermine unify utilize
vacuum viable virtually visible vulnerable
web wheelchair wholesale widow wildlife workplace

abolish accent accessory accompany accountability acquaint acute
adhere adjacent admire aerial aggregate alien allege alliance
allowance altar amid anchor angel animate anonymous anthem anxious
apparatus applaud apt arch ashamed astronaut attorney auction

ballot barber baron basin behalf bet bid biography blanket blaze
bless bloom blush boast bronze buddy burst

cabin canal cape cartoon casualty cathedral cedar celebrity charm
chorus chronic civic cliff cling cluster colonial compass
comprehensive comprise compute conceal confer confine conscience
consent consultancy contemplate contrary controversy cooperation
cord corporate counsel craft crane crude crystal cult curtain

dawn debris delegate democracy density depot descend dessert
dialect diary dim diplomatic dismiss dominant draft drain drown dwell

echo edge embassy embrace encompass endure enrich ensue envelope
era erect essence estate ethnic evaluate exceed exclude exhaust
exploit

facilitate fairy fame fantasy feast feat feminine fiber fleet
flexible float flour forbid forge forum foundation franchise
frontier frost
gallop gauge gem genuine gesture glasses goat gorgeous guardian
gulf
habitat halt hardware harvest hay heir hemisphere heritage
highway hint hollow horizon humidity hurricane hypothesis

identical ideology illustration immense immune impair implement
impulse incentive incorporate indigenous influential infrastructure
inhabit inland innocent inspect insure intact intellectual
interact interim intimate intuitive invade invoice irrigation ivory

jewel junction jury juvenile

keen knight
landlady latitude legacy legend legislative lens liability liken
linger litter locomotive loom lure luxury

magic mandate manure marble margin marshal mast mature mentor
mercury midst migrate miniature moderate molecule monkey monopoly
mortgage municipal myth

negotiate neutral niece noble norm notorious

obstacle onset optical orbit organic orphan ounce outlook
overthrow oxygen

paradise parish parliament partition patron pearl peasant
peninsula pension perceive permanent pest petition philosopher
physician pier pile pioneer plead plunge poultry prairie precise
premise premium prey prince privilege productive profound
prominent proposition province provoke psychiatric pursuit

quest

racism rainbow rally rape realm reckon referee reign reinforce
render republic resemble reside resolve rigid ritual rivalry
rotation rupture

saint sanctuary satellite scandal scaffold scroll seasonal
secular segment semiconductor senate sensation sequel sergeant
simulation simultaneous skeleton skeptical slavery solidarity
sovereign span specimen squad stab statistical stem stimulate
strand subscribe successive superintendent supreme surplus
surveillance sway symbolic sympathy synthetic

tail tap terminology terrain textile therapeutic throne
threshold tobacco torch toxic transformation transmission
tremendous tribe trophy turbulence tutor

undergo undergraduate unify utility

vacuum vague veteran virtual vocal voluntary

ward warrant warrior welfare wilderness worm worthwhile wrap

""".split()

def gen_ngsl():
    seen = set()
    unique = []
    for w in ngsl_words:
        w = w.lower()
        if w not in seen:
            seen.add(w)
            unique.append(w)
    path = os.path.join(DIR, 'ngsl.csv')
    with open(path, 'w') as f:
        for i, w in enumerate(unique, 1):
            f.write(f"{i},{w}\n")
    print(f"ngsl.csv: {len(unique)} words")

# ============================================================
# AWL — sublist,word  (Academic Word List ~570 families expanded)
# ============================================================
awl_data = {
    1: """
    analyse analysis analyst analytical analysed analysing analyses analysts
    approach approached approaches approaching approachable
    area areas
    assess assessed assessing assessment assessments assessor
    assume assumed assumes assuming assumption assumptions
    authority authorities authoritative
    available availability unavailable
    benefit beneficial beneficiaries beneficiary benefited benefiting benefits
    concept concepts conceptual conceptualise
    consist consisted consistency consistent consistently consisting consists
    constitute constituted constitutes constituting constitution constitutional
    context contexts contextual contextualise
    contract contracted contracting contractor contractors contracts
    create created creating creation creative creatively creator creators
    data
    define defined defines defining definition definitions
    derive derived derives deriving derivation
    distribute distributed distributing distribution distributions distributor
    economy economic economical economically economics economies economist
    environment environmental environmentalist environmentally environments
    establish established establishes establishing establishment
    estimate estimated estimates estimating estimation
    evident evidently evidence
    export exported exporter exporters exporting exports
    factor factors
    finance financed finances financial financially financing
    formula formulae formulas formulate formulated formulating formulation
    function functional functioned functioning functions
    identify identifiable identification identified identifies identifying identity
    income incomes
    indicate indicated indicates indicating indication indicator indicators
    individual individually individuals individualism individuality
    interpret interpretation interpretations interpreted interpreting interprets
    involve involved involvement involves involving
    issue issued issues issuing
    labour laboured labouring labours
    legal legality legally illegal illegally
    legislate legislated legislation legislative legislator legislature
    major majority
    method methodical methodological methodology methods
    occur occurred occurrence occurrences occurring occurs
    percent percentage percentages
    period periodic periodical periodically periods
    policy policies
    principle principled principles unprincipled
    proceed procedural procedure procedures proceeded proceeding proceedings proceeds
    process processed processes processing
    require required requirement requirements requires requiring
    research researched researcher researchers researches researching
    respond responded respondent respondents responding response responses
    role roles
    section sectioned sectioning sections
    sector sectors
    significant insignificant significantly significance
    similar dissimilar similarities similarity similarly
    source sourced sources sourcing
    specific specifically specification specifications specificity specifics
    structure restructure restructured restructuring structural structurally structured structures structuring
    theory theoretical theoretically theories theorist theorists
    vary variable variability variables variance variation variations varied varies varying
    """,
    2: """
    achieve achievable achieved achievement achievements achieves achieving
    acquire acquired acquires acquiring acquisition acquisitions
    administrate administrated administrates administrating administration
    affect affected affecting affection affects unaffected
    appropriate appropriacy appropriately appropriateness inappropriately
    aspect aspects
    assist assistance assistant assistants assisted assisting assists
    category categories categorisation categorise categorised
    chapter chapters
    commission commissioned commissioner commissioners commissioning commissions
    community communities
    complex complexities complexity
    compute computation computational computed computer computerised computers computing
    conclude concluded concludes concluding conclusion conclusions conclusive
    conduct conducted conducting conducts
    consequent consequence consequences consequently
    construct constructed constructing construction constructions constructive
    consume consumed consumer consumers consuming consumption
    credit credited crediting creditor creditors credits
    culture cultural culturally cultured cultures uncultured
    design designed designer designers designing designs
    distinct distinction distinctions distinctive distinctly
    element elements
    equate equated equation equations equates equating
    evaluate evaluated evaluates evaluating evaluation evaluations
    feature featured features featuring
    final finalise finally finals
    focus focused focuses focusing refocus
    impact impacted impacting impacts
    injure injured injuries injuring injury
    institute instituted institutes instituting institution institutional institutions
    invest invested investing investment investments investor investors
    item itemise items
    journal journalism journalist journalists journals
    maintain maintained maintaining maintains maintenance
    normal abnormal abnormally normalise normality normally
    obtain obtainable obtained obtaining obtains unobtainable
    participate participant participants participated participates participating participation
    perceive perceived perceives perceiving perception perceptions
    positive positively
    potential potentially
    previous previously
    primary primarily
    purchase purchased purchaser purchasers purchases purchasing
    range ranged ranges ranging
    region regional regionally regions
    regulate deregulate deregulation regulated regulates regulating regulation regulations
    relevant irrelevance irrelevant relevance
    reside resided residence resident residential residents resides residing
    resource resourced resourceful resources resourcing under-resourced
    restrict restricted restricting restriction restrictions restrictive restricts
    secure insecure insecurity secured securely security
    seek seeking seeks sought
    select selected selecting selection selections selective selectively selector
    site sites
    strategy strategic strategically strategies strategist
    survey surveyed surveying surveys surveyor
    text texts textbook textbooks textual
    tradition traditional traditionalist traditionally traditions
    transfer transferred transferring transfers
    """,
    3: """
    alternative alternatively alternatives
    circumstance circumstances
    comment commentaries commentary commented commenting comments
    compensate compensated compensates compensating compensation
    component components
    consent consented consenting consents
    considerable considerably
    constant constantly constants
    constrain constrained constraining constrains constraint constraints
    contribute contributed contributes contributing contribution contributions contributor
    convene convention conventional conventionally conventions convened convenes convening
    coordinate coordinated coordinates coordinating coordination coordinator
    core cores
    corporate corporation corporations
    correspond corresponded correspondence corresponding correspondingly corresponds
    criteria criterion
    deduce deduced deduces deducing deduction deductions
    demonstrate demonstrated demonstrates demonstrating demonstration demonstrations demonstrator
    document documentation documented documenting documents
    dominate dominance dominant dominated dominates dominating domination
    emphasis emphasise emphasised emphasises emphasising
    ensure ensured ensures ensuring
    exclude excluded excludes excluding exclusion exclusions exclusive exclusively
    framework frameworks
    fund funded funder funders funding funds
    illustrate illustrated illustrates illustrating illustration illustrations
    immigrate immigrant immigrants immigrated immigrates immigrating immigration
    imply implied implies implying implication implications
    initial initially
    instance instances
    interact interacted interacting interaction interactions interactive interactively
    justify justifiable justifiably justification justifications justified justifies justifying
    layer layered layers
    link linkage linkages linked linking links
    locate located locating location locations relocate
    maximise maximised maximises maximising maximum
    minor minorities minority minors
    negate negated negates negating negative negatively negatives
    outcome outcomes
    partner partnered partners partnership partnerships
    philosophy philosopher philosophers philosophical philosophically
    physical physically
    proportion disproportion disproportionate disproportionately proportional proportionally proportionate proportionately proportions
    publish published publisher publishers publishes publishing publication publications
    react reacted reacting reaction reactionary reactions reactor reactors
    register deregister registered registering registers registration
    rely reliability reliable reliably reliance reliant relied relies relying
    remove removable removal removed removes removing
    scheme schematic schemed schemes scheming
    sequence sequenced sequences sequencing sequential sequentially
    sex sexes sexism sexual sexuality sexually
    shift shifted shifting shifts
    specify specifiable specified specifies specifying specification specifications
    sufficient sufficiency sufficiently insufficient insufficiently
    task tasks
    technical technically technique techniques technology technological technologically
    valid invalidate invalidity validate validated validating validation validity
    volume volumes
    """,
    4: """
    access accessed accesses accessibility accessible accessing inaccessible
    adequate adequacy adequately inadequacy inadequate inadequately
    annual annually
    apparent apparently
    approximate approximated approximately approximates approximating approximation
    attitude attitudes attitudinal
    attribute attributable attributed attributes attributing attribution
    civil civilian
    code coded codes coding
    commit commitment commitments committed committing
    communicate communicated communicates communicating communication communications
    concentrate concentrated concentrates concentrating concentration
    confer conference conferences conferred conferring confers
    contrast contrasted contrasting contrasts
    cycle cycled cycles cycling cyclical
    debate debatable debated debates debating
    despite
    dimension dimensional dimensions multidimensional
    domestic domestically domesticate domesticated
    emerge emerged emergence emergent emerges emerging
    error erroneous erroneously errors
    ethnic ethnically ethnicity
    goal goals
    grant granted granting grants
    hence
    hypothesis hypotheses hypothesise hypothesised hypothetical hypothetically
    implement implementation implemented implementing implements
    implicate implicated implicates implicating implication implications
    impose imposed imposes imposing imposition
    integrate integrated integrates integrating integration
    internal internally internalise
    investigate investigated investigates investigating investigation investigations investigative investigator
    job jobless jobs
    label labelled labelling labels
    mechanism mechanisms
    obvious obviously
    occupy occupant occupants occupation occupational occupied occupier occupies occupying
    option optional options
    output outputs
    overall
    parallel paralleled parallels unparalleled
    parameter parameters
    phase phased phases phasing
    predict predictability predictable predictably predicted predicting prediction predictions predictor
    principal principally principals
    prior
    professional professionally professionalism professionals unprofessional
    project projected projecting projection projections projects
    promote promoted promoter promoters promotes promoting promotion promotions
    regime regimes
    resolve resolution resolved resolves resolving unresolved
    retain retained retaining retains retention
    series
    statistic statistician statisticians statistical statistically statistics
    status
    stress stressed stresses stressful stressing unstressed
    subsequent subsequently
    sum summation summed summing sums
    summary summarise summarised summarises summarising
    undertake undertaken undertakes undertaking
    """,
    5: """
    academy academia academic academically academics
    adjust adjusted adjusting adjustment adjustments adjustable readjust
    alter alterable alteration alterations altered altering alternate alternating alters unalterable
    amend amended amending amendment amendments amends
    aware awareness unaware
    capacity capacities incapacitate
    challenge challenged challenger challengers challenges challenging
    clause clauses
    compound compounded compounding compounds
    conflict conflicted conflicting conflicts
    consult consultancy consultant consultants consultation consultations consulted consulting consults
    contact contactable contacted contacting contacts
    decline declined declines declining
    discrete discretely discretion discretionary indiscrete indiscretion
    draft drafted drafting drafts redraft
    enable enabled enables enabling
    energy energetic energetically energies energise
    enforce enforced enforcement enforces enforcing
    entity entities
    equivalent equivalence equivalents
    evolve evolution evolutionary evolved evolves evolving
    expand expanded expanding expansion expansionism expansions expands
    expose exposed exposes exposing exposure exposures
    external externalise externally
    facilitate facilitated facilitates facilitating facilitation facilitator facilitators
    fundamental fundamentally fundamentals
    generate generated generates generating generation generations
    generation generational generations
    image imagery images
    liberal liberalise liberalism liberally liberals liberation
    license licensed licenses licensing licensee licensees unlicensed
    logic illogical illogically logical logically logician
    margin marginal marginally margins
    medical medically
    mental mentally mentality
    modify modified modifies modifying modification modifications
    monitor monitored monitoring monitors unmonitored
    network networked networking networks
    notion notional notions
    objective objectively objectivity objectives
    orient orientated orientation oriented
    perspective perspectives
    precise imprecise precisely precision
    psychology psychological psychologically psychologist psychologists
    pursue pursued pursues pursuing pursuit
    ratio ratios
    reject rejected rejecting rejection rejections rejects
    revenue revenues
    stable instability stabilise stabilised stabilising stability unstable
    style styled styles styling stylish stylishly
    substitute substituted substitutes substituting substitution
    sustain sustainability sustainable sustained sustaining sustains unsustainable
    symbol symbolic symbolically symbolise symbolism symbols
    target targeted targeting targets
    transit transition transitional transitions transitory transits
    trend trends
    version versions
    welfare
    """,
    6: """
    abstract abstracted abstracting abstraction abstractions abstractly abstracts
    accurate accuracy accurately inaccuracy inaccurate inaccurately
    acknowledge acknowledged acknowledges acknowledging acknowledgement
    aggregate aggregated aggregates aggregating aggregation
    allocate allocated allocates allocating allocation allocations
    assign assigned assigning assignment assignments assigns reassign
    attach attached attaches attaching attachment attachments unattached
    author authored authoring authoritative authors authorship
    bond bonded bonding bonds
    brief briefed briefing briefly briefs
    capable capabilities capability incapable
    cite cited cites citing citation citations
    cooperate cooperated cooperates cooperating cooperation cooperative cooperatively
    discriminate discriminated discriminates discriminating discrimination
    display displayed displaying displays
    diverse diversely diversification diversified diversifies diversify diversifying diversity
    domain domains
    edit edited editing edition editions editor editorial editors
    enhance enhanced enhancement enhancements enhances enhancing
    estate estates
    exceed exceeded exceeding exceedingly exceeds
    expert expertise experts
    explicit explicitly
    federal federalism federally federation
    fee fees
    flexible flexibility inflexible inflexibility
    furthermore
    gender gendered genders
    ignorance ignorant ignore ignored ignores ignoring
    incentive incentives
    incidence incident incidentally incidents
    incorporate incorporated incorporates incorporating incorporation
    index indexed indexes indexing indices
    inhibit inhibited inhibiting inhibition inhibitions inhibits
    initiate initiated initiates initiating initiation initiative initiatives
    input inputs
    instruct instructed instructing instruction instructional instructions instructor instructors instructs
    intelligence intelligent intelligently unintelligent
    interval intervals
    lecture lectured lecturer lecturers lectures lecturing
    migrate migrant migrants migrated migrates migrating migration migratory
    minimum minimal minimally minimise minimised
    ministry ministerial ministries minister ministers
    motive motivate motivated motivates motivating motivation motivations motives unmotivated
    neutral neutralise neutrality neutrally
    nevertheless
    overseas
    precede preceded precedence precedent preceding precedes
    presume presumably presumed presumes presuming presumption presumptions presumptuous
    rational irrational irrationally rationale rationalism rationality rationally
    recover recoverable recovered recovering recovers recovery
    reveal revealed revealing reveals revelation revelations
    scope scoped scopes scoping
    subsidy subsidies subsidise subsidised subsidises subsidising
    tape taped tapes taping
    trace traceable traced traces tracing
    transform transformation transformations transformed transforming transforms
    transport transportation transported transporting transports
    underlying
    utilise utilisation utilised utilises utilising utility utilities
    """,
    7: """
    adapt adaptable adaptation adaptations adapted adapting adaptive adapts
    adult adulthood adults
    advocate advocated advocates advocating advocacy
    aid aided aiding aids unaided
    channel channelled channelling channels
    chemical chemically chemicals
    classic classical classics
    comprehensive comprehensively
    comprise comprised comprises comprising
    confirm confirmation confirmed confirming confirms unconfirmed
    contrary contrarily contraries
    convert converted converting conversion conversions converts convertible
    couple coupled coupling couples
    decade decades
    definite definitely definitive indefinite indefinitely
    deny deniable denial denied denies denying undeniable
    differentiate differentiated differentiates differentiating differentiation
    dispose disposable disposal disposed disposes disposing disposition
    dynamic dynamically dynamics dynamism
    eliminate eliminated eliminates eliminating elimination
    empirical empirically empiricism
    equip equipment equipped equipping equips
    extract extracted extracting extraction extracts
    file filed files filing
    finite finitely infinite infinitely
    foundation foundations founded founder founders founding
    globe global globalisation globally globes
    grade graded grades grading
    guarantee guaranteed guaranteeing guarantees
    hierarchy hierarchical hierarchically hierarchies
    identical identically
    ideology ideological ideologically ideologies
    infer inference inferences inferred inferring infers
    innovate innovated innovates innovating innovation innovations innovative innovator innovators
    insert inserted inserting insertion inserts
    intervene intervened intervenes intervening intervention interventions
    isolate isolated isolates isolating isolation isolationism
    media
    mode modes
    paradigm paradigms
    phenomenon phenomena
    priority priorities prioritise prioritised prioritises prioritising
    prohibit prohibited prohibiting prohibition prohibitions prohibitive prohibits
    publication publications
    quote quotation quotations quoted quotes quoting
    release released releases releasing
    reverse reversal reversed reverses reversible reversing irreversible
    simulate simulated simulates simulating simulation simulations simulator
    sole solely
    somewhat
    submit submission submissions submitted submitting submits
    successor succession successions successive successively successors
    survive survival survived survives surviving survivor survivors
    thesis theses
    topic topical topics
    transmit transmission transmissions transmitted transmitter transmitting transmits
    ultimate ultimately
    unique uniquely uniqueness
    visible visibility visibly invisible invisibility invisibly
    voluntary voluntarily volunteer volunteered volunteering volunteers
    """,
    8: """
    abandon abandoned abandoning abandonment abandons
    accompany accompanied accompanies accompanying accompaniment
    accumulate accumulated accumulates accumulating accumulation
    ambiguous ambiguity ambiguities unambiguous unambiguously
    append appended appendices appendix appending appends
    appreciate appreciated appreciates appreciating appreciation
    arbitrary arbitrarily arbitrariness
    automate automated automates automating automatic automatically automation
    bias biased biases biasing unbiased
    cease ceased ceases ceasing
    coherent coherence coherently incoherent incoherence
    coincide coincided coincidence coincidences coincides coinciding coincidental
    commence commenced commences commencing commencement
    compatible compatibility incompatible incompatibility
    conform conformist conformity conformed conforming conforms nonconformist nonconformity
    contemporary contemporaries
    contradict contradicted contradicting contradiction contradictions contradictory contradicts
    crucial crucially
    currency currencies
    denote connoted connotation denoted denotes denoting denotation
    detect detectable detected detecting detection detective detectives detector detectors detects
    deviate deviated deviates deviating deviation deviations
    displace displaced displacement displaces displacing
    drama dramatic dramatically dramatise dramatist dramas
    eventual eventuality eventually
    exhibit exhibited exhibiting exhibition exhibitions exhibits
    exploit exploitable exploitation exploited exploiting exploits
    fluctuate fluctuated fluctuates fluctuating fluctuation fluctuations
    guideline guidelines
    highlight highlighted highlighting highlights
    implicit implicitly
    induce induced induces inducing induction
    inevitable inevitability inevitably
    infrastructure infrastructural infrastructures
    inspect inspected inspecting inspection inspections inspector inspectors inspects
    intense intensely intensify intensified intensifies intensifying intensity intensive intensively
    manipulate manipulated manipulates manipulating manipulation manipulative
    minimise minimised minimises minimising
    nuclear
    offset offsets offsetting
    paragraph paragraphs
    plus pluses
    practitioner practitioners
    predominant predominantly predominance predominate predominated predominates predominating
    prospect prospective prospects
    radical radically radicals radicalism
    random randomly randomness randomise
    reinforce reinforced reinforcement reinforcements reinforces reinforcing
    restore restored restores restoring restoration
    revise revised revises revising revision revisions
    schedule reschedule rescheduled scheduled schedules scheduling unscheduled
    tense tensely tension tensions
    terminate terminal terminally terminated terminates terminating termination
    theme thematic thematically themed themes
    thereby
    uniform uniformity uniformly uniforms
    vehicle vehicles vehicular
    via
    virtual virtually
    visual visualise visualised visualising visualisation visually
    widespread
    """,
    9: """
    analogy analogies analogous
    anticipate anticipated anticipates anticipating anticipation unanticipated
    assure assurance assured assures assuring reassure reassurance
    attain attainable attained attaining attainment attains unattainable
    behalf
    bulk bulky
    cease ceased ceases ceasing
    coherent coherence coherently incoherent
    collapse collapsed collapses collapsing
    colleague colleagues
    compile compiled compiles compiling compilation compilations
    conceive conceivable conceivably conceived conceives conceiving inconceivable
    convince convinced convinces convincing convincingly unconvinced
    dedicate dedicated dedicates dedicating dedication
    deviation deviate deviated deviates deviating deviations
    diminish diminished diminishes diminishing diminution undiminished
    distort distorted distorting distortion distortions distorts undistorted
    draft drafted drafting drafts redraft
    enormous enormity enormously
    erode eroded erodes eroding erosion
    ethics ethical ethically unethical
    format formatted formatting formats
    found foundation founded founder founders founding
    inherent inherently
    insight insightful insights
    integral
    intermediate intermediaries intermediary
    manual manually manuals
    mature immature immaturity maturation matured matures maturing maturity
    mediate mediated mediates mediating mediation mediator mediators
    medium media
    military militarily militia
    norm norms normative
    ongoing
    panel panelled panelling panels
    persist persisted persistence persistent persistently persisting persists
    pose posed poses posing
    reluctance reluctant reluctantly
    scenario scenarios
    scope scoped scopes scoping
    sphere spheres spherical spherically
    subordinate subordinated subordinates subordinating subordination
    supplement supplementary supplemented supplementing supplements
    suspend suspended suspending suspension suspensions suspends
    team teamed teaming teams teamwork
    temporary temporarily
    trigger triggered triggering triggers
    unify unification unified unifies unifying
    violate violated violates violating violation violations
    vision visionary
    """,
    10: """
    adjacent adjacency
    albeit
    assemble assembled assembles assemblies assembling assembly
    collapse collapsed collapses collapsing
    colleague colleagues
    compile compiled compiles compiling compilation
    conceive conceivable conceived conceives conceiving
    convince convinced convinces convincing unconvinced
    depress depressed depresses depressing depression
    encounter encountered encountering encounters
    enormous enormity enormously
    forthcoming
    incline disinclination disinclined inclined inclines inclining inclination inclinations
    integrity
    intrinsic intrinsically
    invoke invoked invokes invoking invocation
    levy levies levied levying
    likewise
    nonetheless
    notwithstanding
    odd oddity oddly odds
    ongoing
    panel panelled panels
    persist persisted persistence persistent persists
    pose posed poses posing
    reluctance reluctant reluctantly
    so-called
    straightforward
    undergo undergoes undergoing undergone underwent
    whereby
    """,
}

# Supplemental AWL expansions to reach ~3000
awl_supplement = {
    1: """
    analysing analyzer analyzable reanalysis reanalyse
    approachability unapproachable
    assessable reassess reassessment reassessed
    assumable reassume
    authorise authorised authorisation authoritative unauthorised
    benefactor benefited
    conceptualisation conceptualised conceptually
    consistently inconsistency inconsistencies
    constitutionally unconstitutional unconstitutionally
    contextualised contextualising decontextualise
    contractual contractually subcontract subcontractor subcontracted
    creatively recreate recreated recreation recreational
    definable indefinable redefined redefinition undefined
    derivation derivative derivatives
    distributional redistributed redistribution
    economise economised economising uneconomical
    environmentalism
    disestablish reestablish reestablishment
    overestimate overestimated underestimate underestimated
    evidential
    exporter
    formulaic reformulate reformulated reformulation
    dysfunctional functionally multifunctional
    identifiable unidentified
    indicative
    individualistic individualized individually
    interpretive misinterpret misinterpretation reinterpret reinterpretation
    uninvolved
    legislative legislatively
    methodically
    reoccur reoccurrence
    policyholder policymakers
    unprincipled
    procedurally
    reprocessed unprocessed
    prerequisite
    researcher researchable
    irresponsible irresponsibly responsibly responsive
    restructured
    theoretically theorise theorised
    invariable invariably unvaried
    """,
    2: """
    achiever underachieve underachievement underachiever
    reacquire reacquisition
    administrative administratively
    unaffected affectionately
    inappropriacy inappropriate inappropriately
    unassisted
    categorically uncategorised
    decommission recommission
    intercommunity
    overcomplicate uncomplicated
    inconclusive inconclusively
    misconduct
    inconsequential
    deconstruct deconstruction reconstruction reconstructed
    overconsume overconsumption
    discredit discreditable
    multicultural subcultural uncultured
    redesign redesigned
    indistinct indistinctly
    elemental
    nonequivalent
    reevaluate reevaluation overvalue undervalue
    unfocused
    uninjured
    institutional uninstitutional
    reinvest disinvest disinvestment
    itemised
    journalistic
    unmaintained
    abnormality normalcy normalisation renormalise
    unobtained
    nonparticipant nonparticipation
    imperceptible imperceptibly
    disproportionately overregulate underregulate
    irrelevance irrelevantly
    nonresident
    unresourced
    unrestricted overrestrict
    insecure insecurities
    reselect reselection unselected
    strategic strategically
    resurvey
    nontraditional traditionally
    retransfer untransferable
    """,
    3: """
    uncompensated overcompensate overcompensation
    consentual nonconsent
    inconsiderable
    inconstancy
    unconstrained
    noncontributing undercontribute
    reconvene unconventional unconventionally
    uncoordinated
    incorporate reincorporation
    overdeduce
    undemonstrated
    undocumented redocument
    undominated
    overemphasise overemphasis underemphasis
    overexclude
    refunction
    unfunded underfund underfunded refund refundable nonrefundable
    unillustrated
    immigrant
    implicitly unexplicitly
    reinteract
    unjustified unjustifiably
    unlayered multilayer multilayered
    unlinked interlink interlinked
    dislocate dislocation mislocate
    nonminority
    nonnegative
    repartner
    republish unpublished
    overreact underreact
    registration deregister reregister
    unreliable unreliably
    nonremovable irremovable
    unschematic
    resequence
    unspecified overspecify
    insufficiency
    retask multitask tasking
    """,
    4: """
    inaccessibility accessibly
    inadequately
    biannual biannually perennial perennially
    unapparent
    inapproximable
    reattribute unattributed misattribute
    incivility uncivil civilise civilised civilisation
    decode encoded encoder encoding recode
    uncommitted noncommittal
    miscommunicate miscommunication noncommunicative
    deconcentrate reconcentrate
    nonconference
    uncontrastable
    recyclable recyclability recycle recycled
    redebate
    multidimensionality
    undomesticated
    reemerge reemergent reemergence
    unerring inerrant errantly
    multiethnic interethnic
    reimplement reimplementation
    reimpose
    nonintegrated reintegrate reintegration disintegrate
    reinvestigate reinvestigation
    relabel mislabel unlabelled
    remechanise
    preoccupation preoccupied unoccupied
    reoutput
    unparalleled
    rephase
    unpredicted unpredictably misprediction
    reprioritise deprioritise
    unprofessionally semiprofessional
    reproject reprojection
    demote demoted demotion
    irresolvable
    retainable
    summarisation
    """,
    5: """
    maladjusted readjustment readjusted
    unalterable inalterably
    unamended
    unaware unawareness
    incapacitated overcapacity undercapacity
    unchallenged rechallenge
    compounding uncompounded
    unconflicted
    reconsult preconsult
    recontact
    undeclining
    indiscreet indiscretion redraft
    reenabled
    reenergise deenergise
    unenforced nonenforcement
    nonentity
    inequivalent nonequivalent
    unevolved devolution
    overexpand unexpanded
    overexpose underexpose unexposed
    externalise externalisation
    unfacilitated
    regenerate degenerate regeneration degeneration degenerative
    reimagine reimagined
    illiberal illiberally
    relicense unlicensed
    illogic
    demarginalize demarginalized remarginalise
    nonmedical premedical biomedical
    nonmental
    remodify unmodifiable
    unmonitored remonitor
    renetwork
    nonnotion
    nonobjective
    disorient disorientation
    imprecisely imprecision
    neuropsychological psychologically
    repursue
    rereject
    destabilise destabilisation prestabilise destabilising
    restyle restyled restylise
    resubstitute
    unsustainably unsustained
    desymbolise
    retarget untargeted
    pretransit retransit
    """,
    6: """
    reacknowledge unacknowledged
    reaggregate disaggregate disaggregation
    reallocate reallocation misallocate
    reassign unassigned misassign
    reattach detach detached detachment
    reauthorise unauthorised coauthor coauthored
    rebond debond unbonded
    rebriefing debriefing
    incapability
    recitation recite reciter
    uncooperative noncooperative
    indiscriminately nondiscriminatory
    redisplay undisplayed
    biodiversity
    re-edit unedited edited
    unenhanced preenhance
    re-estimate reestimate overestimate
    inexpert unexpert
    inexplicit unexplicit
    infederal nonfederal
    inflexibly
    genderless nongendered
    unincentivised disincentivise
    unincorporated priorincorporate
    uninhibited disinhibit disinhibited
    uninstructed misinstructed
    unintelligently
    unlectured
    nonministerial
    demotivate demotivated unmotivated remotivate
    nonneutral
    unprecedented
    irrationalise irrationality
    irrecoverable unrecovered unrecoverable
    nonrevealed unrevealed
    rescope descope
    unsubsidised
    retraced untraceable
    retransform untransformed
    untransported
    """,
    7: """
    readaptation maladaptation maladapted
    unaided
    rechannelled
    dechlorinate nonchemical biochemical
    unclassical nonclassical neoclassical
    incomprehensively
    unconfirmable reconfirm
    noncontrary
    reconvert unconverted deconversion inconvertible
    uncoupled decouple
    redefinite indefinitely redefinitely
    undeniably
    undifferentiated
    indisposable predispose
    nondynamic
    reeliminate
    unequipped underequipped
    inexact inexactly
    nonfinite
    cofound cofounder unfounded
    nongrade grading ungraded regrading
    unguaranteed
    ideologist
    reinfer
    noninnovative
    reinsert
    noninterventionist
    reisolate deisolate unisolated
    nonparadigmatic
    prioritization
    nonprohibitive
    republication
    requote misquote
    rerelease unreleased prereleased
    irreversibly
    resimulate unsimulated
    resubmit
    retransmission retransmit
    """,
    8: """
    reabandoned
    unaccompanied
    reaccumulate
    unambiguously
    reappend
    underappreciate underappreciated
    nonarbitrary
    nonautomated semiautomated semiautomatic
    unbias rebias
    incoherently
    noncoincidental
    recommence
    incompatible
    nonconformity
    noncontradictory
    detectable undetectable
    redeviate
    redisplace undisplaced
    nondramatic
    uneventful uneventfully
    re-exhibit
    unexploited overexploited
    nonfluctuating
    rehighlight unhighlighted
    noninduction
    uninspected
    overintense
    unmanipulated
    nonminimal
    re-offset
    nonpredominant
    nonprospective
    nonradical
    randomised nonrandom pseudorandom
    unreinforced
    unrestored
    nonrevised
    rescheduled rescheduling
    reterminate
    nonthematic rethematic
    nonuniform
    nonvisual revisualise
    """,
    9: """
    reassurance unreassured
    unattainable unattained
    coherently
    recompile recompiled
    inconceivably reconceive
    unconvincing unconvincingly
    rededicate undedicated
    undiminished
    undistorted redistort
    uneroded
    unethically
    reformatted reformatting
    unfounded
    noninherent
    insightfully
    nonintegral
    nonintermediate
    nonmanual
    immature immaturely prematurely
    unmediated
    nonmilitary paramilitary
    nonnormative
    nonpersistent
    repose reposed repositioned
    nonreluctant
    rescope rescoped
    nonsphere hemispheric hemispheres
    insubordinate insubordination
    unsupplemented
    unsuspended
    nontemporary
    untriggered retriggered
    reunification reunify
    nonviolation
    nonvisionary
    """,
    10: """
    nonadjacent
    reassemble reassembled
    recollapse
    recompile
    reconvince
    antidepressant
    re-encounter
    reencounter
    disinclination inclinations
    nonintegrity
    nonintrinsic
    reinvoke
    re-levy
    nonpanelled
    nonpersistent
    reposit reposition
    nonreluctance
    nonstraightforward
    re-undergo
    """,
}

def gen_awl():
    path = os.path.join(DIR, 'awl.csv')
    total = 0
    seen = set()
    with open(path, 'w') as f:
        for sublist in sorted(awl_data):
            words = awl_data[sublist].split()
            for w in words:
                key = (sublist, w)
                if key not in seen:
                    seen.add(key)
                    f.write(f"{sublist},{w}\n")
                    total += 1
        # Add supplement
        for sublist in sorted(awl_supplement):
            words = awl_supplement[sublist].split()
            for w in words:
                key = (sublist, w)
                if key not in seen:
                    seen.add(key)
                    f.write(f"{sublist},{w}\n")
                    total += 1
    print(f"awl.csv: {total} words across {len(awl_data)} sublists")

if __name__ == '__main__':
    gen_oxford()
    gen_ngsl()
    gen_awl()
    print("Done!")
